extern crate hyper;

use std::env;
use std::time::Duration;
use std::thread::sleep;
use std::process::exit;
use hyper::Client;
use hyper::header::Connection;

fn check_host(host: &str, attempt: i8) {
  if attempt > 5 {
    println!("Couldn't reach host {} after {} attempts", host, attempt);
    exit(1);
  }

  let client = Client::new();

  let result = client
    .get(host)
    .send();

  match result {
    Ok(val) => {
      println!("Host ok {}", host);
      exit(0);
    }
    Err(e) => {
      println!("Host not found {} - {}", host, e);
      let dur = Duration::new(1, 0);
      sleep(dur);
      check_host(host, attempt + 1);
    },
  }
}

fn main() {
  let args: Vec<String> = env::args().collect();

  match args.get(1) {
    Some(host) => {
      println!("Looking for {}", host);
      check_host(&host, 1);
    },
    None => {
      println!("Expected host");
      exit(1);
    }
  }
}
