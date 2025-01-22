use std::sync::mpsc;
use std::thread;
use std::time::Duration;

// Simulates incoming raw data
fn generate_data() -> f64 {
    let raw_value = rand::random::<f64>() * 100.0;
    raw_value
}

// Processes the data
fn process_data(raw_value: f64) -> f64 {
    raw_value * 1.1 // Example: apply a multiplier for processing
}

fn main() {
    // Create a channel for communication between threads
    let (tx, rx) = mpsc::channel();

    // Spawn a thread to generate raw data
    thread::spawn(move || {
        loop {
            let raw_value = generate_data();
            println!("Generated raw value: {:.2}", raw_value);
            if let Err(err) = tx.send(raw_value) {
                eprintln!("Error sending data: {}", err);
                break;
            }
            thread::sleep(Duration::from_secs(1)); // Simulate real-time data generation
        }
    });

    // Main thread processes the incoming data
    println!("Starting data processing...");
    for received in rx {
        let processed_value = process_data(received);
        println!(
            "Received raw value: {:.2} -> Processed value: {:.2}",
            received, processed_value
        );
    }
}
