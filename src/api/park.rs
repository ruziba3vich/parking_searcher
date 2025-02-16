use actix_web::{get, post, put, delete, web, HttpResponse, Responder};
use diesel::PgConnection;
use uuid::Uuid;
use crate::repositories::park_repo::ParkRepo;
use crate::models::park::Park;

#[post("/parks")]
async fn create_park(conn: web::Data<PgConnection>, new_park: web::Json<Park>) -> impl Responder {
    match ParkRepo::create(&mut conn.get().unwrap(), &new_park.into_inner()) {
        Ok(park) => HttpResponse::Created().json(park),
        Err(_) => HttpResponse::InternalServerError().finish(),
    }
}

#[get("/parks")]
async fn get_all_parks(conn: web::Data<PgConnection>) -> impl Responder {
    match ParkRepo::get_all(&mut conn.get().unwrap()) {
        Ok(parks) => HttpResponse::Ok().json(parks),
        Err(_) => HttpResponse::InternalServerError().finish(),
    }
}

#[get("/parks/{id}")]
async fn get_park_by_id(conn: web::Data<PgConnection>, park_id: web::Path<Uuid>) -> impl Responder {
    match ParkRepo::get_by_id(&mut conn.get().unwrap(), park_id.into_inner()) {
        Ok(park) => HttpResponse::Ok().json(park),
        Err(_) => HttpResponse::NotFound().finish(),
    }
}

#[put("/parks/{id}")]
async fn update_park(conn: web::Data<PgConnection>, park_id: web::Path<Uuid>, park_data: web::Json<Park>) -> impl Responder {
    match ParkRepo::update(&mut conn.get().unwrap(), park_id.into_inner(), &park_data.into_inner()) {
        Ok(park) => HttpResponse::Ok().json(park),
        Err(_) => HttpResponse::InternalServerError().finish(),
    }
}

#[delete("/parks/{id}")]
async fn delete_park(conn: web::Data<PgConnection>, park_id: web::Path<Uuid>) -> impl Responder {
    match ParkRepo::delete(&mut conn.get().unwrap(), park_id.into_inner()) {
        Ok(_) => HttpResponse::NoContent().finish(),
        Err(_) => HttpResponse::InternalServerError().finish(),
    }
}
