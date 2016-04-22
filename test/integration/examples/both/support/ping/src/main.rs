extern crate hyper;

use std::env;
use std::time::Duration;
use std::thread::sleep;
use std::process::exit;
use hyper::Client;

fn check_host(host: &str) {

  let mut attempt = 1;
  let client = Client::new();

  while attempt < 6 {
    attempt += 1;

    let result = client
      .get(host)
      .send();

    match result {
      Ok(_) => {
        println!("Host ok {}", host);
        exit(0);
      }
      Err(e) => {
        println!("Host not found {} - {}", host, e);
        let dur = Duration::new(1, 0);
        sleep(dur);
      },
    }
  }
  
  println!("Couldn't reach host {} after {} attempts", host, attempt);
  exit(1);

}

fn main() {
  let args: Vec<String> = env::args().collect();

  match args.get(1) {
    Some(host) => {
      println!("Looking for {}", host);
      check_host(&host);
    },
    None => {
      println!("Expected host parameter");
      exit(1);
    }
  }
}
