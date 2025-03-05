use std::{ffi::OsStr, os::unix::ffi::OsStrExt, path::PathBuf};

use bitcode::{Encode, Decode};

use super::{ScriptDBImpl, ScriptDB};

#[derive(Encode, Decode, Debug, Clone, Default)]
pub(crate) struct ScriptDBv1 {
    files: Vec<Vec<u8>>,
    repos: Vec<Vec<u8>>,
}

impl ScriptDBImpl for ScriptDBv1 {
    const VERSION: u16 = 1;

    fn files(&self) -> Vec<PathBuf> {
        Self::bytes_to_pathbuf(&self.files)
    }

    fn repos(&self) -> Vec<PathBuf> {
        Self::bytes_to_pathbuf(&self.repos)
    }

    fn encode(&self) -> Vec<u8> {
        bitcode::encode(self)
    }

    fn decode(contents: &[u8]) -> Self {
        bitcode::decode(contents).unwrap()
    }
}

impl ScriptDBv1 {
    fn bytes_to_pathbuf(array: &[Vec<u8>]) -> Vec<PathBuf> {
        let mut out = vec![];

        for element in array.iter() {
            out.push(PathBuf::from(OsStr::from_bytes(element)));
        }

        out
    }
}

impl ScriptDBv1 {
    pub fn from_script_db(script_db: &ScriptDB) -> Self {
        let mut files = vec![];
        let mut repos = vec![];

        for file in script_db.files.iter() {
            files.push(file.as_os_str().as_bytes().to_owned())
        }

        for repo in script_db.repos.iter() {
            repos.push(repo.as_os_str().as_bytes().to_owned())
        }

        Self { files, repos }
    }
}
