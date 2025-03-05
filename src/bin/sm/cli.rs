use std::path::PathBuf;

use clap::Parser;

#[derive(Parser, Clone, Debug)]
pub struct Cli {
    #[arg()]
    pub name: Option<String>,

    /// Add a local file or folder to the script store
    #[arg(short, long, value_parser, num_args = 1.., value_delimiter = ' ')]
    pub add: Option<Vec<PathBuf>>,

    /// Remove scripts and repos from the store
    #[arg(short = 'x', long, value_parser, num_args = 1.., value_delimiter = ' ')]
    pub remove: Option<Vec<String>>,

    // /// Add a repo to the script store from a git repo
    // #[arg(short = 'g', long, value_parser, num_args = 1.., value_delimiter = ' ')]
    // pub add_git: Option<String>,

    /// List the scripts in the store
    #[arg(short, long)]
    pub list: bool,

    /// Edit a script
    #[arg(short, long)]
    pub edit: Option<String>,
}

impl Cli {
    pub fn has_operation(&self) -> bool {
        self.add.is_some() || self.remove.is_some() || self.list /* || self.add_git.is_some() */|| self.edit.is_some()
    }
}
