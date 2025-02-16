mod models;

use axum::{routing::{get, post}, Router, Json, extract::Path};
use std::{net::SocketAddr, sync::Arc};
use tokio::sync::Mutex;
use tokio::net::TcpListener;
use tower_http::trace::TraceLayer;
use tracing_subscriber;
use models::Park;

type ParkDB = Arc<Mutex<Vec<Park>>>;

#[tokio::main]
async fn main() {
    tracing_subscriber::fmt()
        .with_env_filter("trace")
        .init();

    let parks: ParkDB = Arc::new(Mutex::new(vec![
        Park { id: 1, name: "Central Park".into(), location: "New York".into() },
        Park { id: 2, name: "Hyde Park".into(), location: "London".into() },
    ]));

    let app = Router::new()
        .route("/", get(|| async { "Hello, world!" }))
        .route("/parks", get(get_parks).post(add_park))
        .route("/parks/:id", get(get_park_by_id))
        .layer(TraceLayer::new_for_http())
        .with_state(parks);

    let addr: SocketAddr = "127.0.0.1:3000".parse().expect("Failed to parse socket address");
    let listener = TcpListener::bind(addr).await.expect("Failed to bind to address");

    println!("🚀 Server running at http://{}", addr);

    axum::serve(listener, app).await.expect("Server failed");
}

async fn get_parks(state: axum::extract::State<ParkDB>) -> Json<Vec<Park>> {
    let parks = state.lock().await;
    Json(parks.clone())
}

async fn get_park_by_id(Path(id): Path<u32>, state: axum::extract::State<ParkDB>) -> Json<Option<Park>> {
    let parks = state.lock().await;
    Json(parks.iter().find(|p| p.id == id).cloned())
}

async fn add_park(state: axum::extract::State<ParkDB>, Json(new_park): Json<Park>) -> Json<String> {
    let mut parks = state.lock().await;
    parks.push(new_park);
    Json("Park added successfully".to_string())
}
