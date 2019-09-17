extern crata regex;
use std::io;
use  regex::Regex;
fn main(){
 let mut input = string::new();
 io::stdln().read_line(&mut input)?;
   ok()
 let re=Regex::new(r"^\w+@[[:word:]]+\.[[:word:]]$").unwrap();
    match re.capture(input) {
Some(caps) =>println!("valid mail"),
None =>println!("invalid mail")
}
}