use std::io::prelude::*;
use std::net;

async fn cheapo_request(host: &str, port: u16, path: &str) -> std::io::Result<String> {
    let mut socket = net::TcpStream::connect((host, port))?;

    let request = format!("GET {} HTTP/1.1\r\nHost: {}\r\n\r\n", path, host);
    socket.write_all(request.as_bytes())?;
    socket.shutdown(net::Shutdown::Write)?;

    let mut response = String::new();
    socket.read_to_string(&mut response)?;

    Ok(response)
}

fn main() -> std::io::Result<()> {
    use async_std::task;

    let response = task::block_on(async { cheapo_request("example.com", 80, "/").await })?;
    println!("{}", response);
    Ok(())
}
