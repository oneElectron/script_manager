mod script_db_v1;
use script_db_v1::ScriptDBv1;

use std::{
    ffi::OsStr,
    os::unix::fs::PermissionsExt,
    path::{Path, PathBuf},
};

type CurrentDBVersion = ScriptDBv1;

const SCRIPTDB_MAGIC_BYTES: &[u8; 8] = b"scriptdb";

const SCRIPT_TEMPLATE: &str = include_str!("default_script.sh");

trait ScriptDBImpl: Default {
    const VERSION: u16;
    fn files(&self) -> Vec<PathBuf>;
    fn repos(&self) -> Vec<PathBuf>;

    fn encode(&self) -> Vec<u8>;
    fn decode(contents: &[u8]) -> Self;

    fn write_to_disk<P: AsRef<Path>>(&self, p: P) {
        let mut contents = self.encode();
        for b in Self::VERSION.to_le_bytes().into_iter().rev() {
            contents.insert(0, b);
        }
        for (i, b) in SCRIPTDB_MAGIC_BYTES.iter().enumerate() {
            contents.insert(i, *b);
        }

        std::fs::write(p, contents).expect("Unable to write to database file");
    }

    fn to_script_db(&self) -> ScriptDB {
        let path = xdg::BaseDirectories::new()
            .unwrap()
            .get_data_file("script_manager");

        debug_assert!(path.is_absolute());

        ScriptDB {
            files: self.files(),
            repos: self.repos(),
            file_write_cache: None,
            path,
        }
    }
}

pub struct ScriptDB {
    files: Vec<PathBuf>,
    repos: Vec<PathBuf>,
    file_write_cache: Option<Vec<(PathBuf, Vec<u8>)>>,
    path: PathBuf,
}

impl ScriptDB {
    /// Read the script database from the disk, creates a new database if none exists
    pub fn read_from_disk() -> ScriptDB {
        let path = xdg::BaseDirectories::new()
            .expect("Unable to find xdg data dir")
            .get_data_file("script_manager/database.scriptdb");

        let contents = std::fs::read(&path);

        if contents.is_err() || contents.as_ref().unwrap().is_empty() {
            return Self::new();
        }

        let contents = contents.unwrap();

        if *SCRIPTDB_MAGIC_BYTES != contents[0..8] {
            panic!("The magic bytes are incorrect");
        }

        let version = u16::from_le_bytes(contents[8..10].try_into().unwrap());

        decode_version(version, &contents[10..])
    }

    pub fn list(&self) -> Vec<PathBuf> {
        let mut out = vec![];

        for f in self.files.iter() {
            out.push(self.path.join(f));
        }

        out
    }

    /// Add a script to the script store.
    /// Will overwrite the script if it exists.
    pub fn add_script<S: AsRef<OsStr>>(
        &mut self,
        name: S,
        contents: Vec<u8>,
    ) -> Result<(), AddScriptError> {
        if self.file_write_cache.is_none() {
            self.file_write_cache = Some(Default::default())
        }

        self.file_write_cache
            .as_mut()
            .unwrap()
            .push((self.path.join("custom").join(name.as_ref()), contents));

        Ok(())
    }

    pub fn edit_script<S: AsRef<OsStr>>(&mut self, name: S) {
        let file = self.get_script(name.as_ref());

        let contents = if let Some(s) = file {
            edit::edit(std::fs::read_to_string(s).unwrap()).unwrap()
        } else {
            edit::edit(SCRIPT_TEMPLATE).unwrap()
        };

        self.add_script(name.as_ref(), contents.bytes().collect())
            .unwrap();
    }

    pub fn write(&mut self) {
        if self.file_write_cache.is_none() {
            return;
        }

        let p = self.path.join("custom");
        if !p.exists() {
            std::fs::create_dir_all(&p).unwrap();
        } else if !p.is_dir() {
            panic!();
        }

        for (path, contents) in self.file_write_cache.as_ref().unwrap() {
            if std::fs::write(path, contents).is_ok() {
                if !self.files.contains(&path.strip_prefix(&self.path).unwrap().to_owned()) {
                    self.files
                        .push(path.strip_prefix(&self.path).unwrap().to_owned());
                }

                std::fs::set_permissions(path, std::fs::Permissions::from_mode(0o755)).unwrap();
            } else {
                println!(
                    "Unknown error occured while writing file: {}",
                    path.display()
                );
            }
        }

        self.sort_files();

        // Make sure the xdg data path exists
        if !self.path.exists() {
            std::fs::create_dir_all(&self.path).unwrap();
        }

        CurrentDBVersion::from_script_db(self).write_to_disk(self.path.join("database.scriptdb"));
    }

    pub fn get_script<S: AsRef<OsStr>>(&self, name: S) -> Option<PathBuf> {
        let r = self
            .files
            .binary_search_by(|x| x.file_name().unwrap().cmp(name.as_ref()));

        if let Ok(s) = r {
            return Some(self.path.join(&self.files[s]));
        }

        None
    }

    fn new() -> ScriptDB {
        CurrentDBVersion::default().to_script_db()
    }

    fn sort_files(&mut self) {
        self.files
            .sort_by(|x, y| x.file_name().unwrap().cmp(y.file_name().unwrap()));
    }

    // pub fn add_repo(url: &str) {
    //     todo!();
    // }
}

impl Drop for ScriptDB {
    fn drop(&mut self) {
        self.write();
    }
}

fn decode_version(version: u16, contents: &[u8]) -> ScriptDB {
    match version {
        1 => ScriptDBv1::decode(contents),
        _ => panic!("Incorrect version of the data base, maybe you used a newer version of script manager previously?"),
    }.to_script_db()
}

#[derive(Debug)]
pub enum AddScriptError {
    /// The added script already exists
    AlreadyExists,
}
