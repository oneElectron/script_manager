mod cli;
use std::{ffi::OsString, process::Stdio};

use cli::Cli;

use clap::Parser;

fn main() {
    let args = Cli::parse();

    if !args.has_operation() && args.name.is_none() {
        println!("You must give an operation or the name of a script");
        return;
    }

    let mut db = script_manager::ScriptDB::read_from_disk();

    if let Some(name) = args.name {
        let script_path = db.get_script(&name).unwrap();

        let mut env_args = std::env::args_os();
        let os_name = OsString::from(&name);
        let i = env_args.position(|x| x == os_name).unwrap();

        let c = std::process::Command::new(script_path).args(
            std::env::args_os().skip(i),
        ).stdin(Stdio::inherit()).stdout(Stdio::inherit()).stderr(Stdio::inherit()).spawn();

        c.unwrap().wait().unwrap();

        return; // Run the script and exit, ignoring any other options
    }

    if args.list {
        for script in db.list() {
            println!("{}", script.file_name().unwrap().to_str().unwrap());
        }
    }

    if let Some(files) = args.add {
        let mut contents = vec![];

        for file in files {
            contents.push((
                file.file_name().unwrap().to_str().unwrap().to_owned(),
                std::fs::read_to_string(file).unwrap(),
            ));
        }

        for (name, content) in contents {
            db.add_script(name, content.bytes().collect()).unwrap();
        }
    }

    if let Some(script) = args.edit {
        db.edit_script(script);
    }
}
