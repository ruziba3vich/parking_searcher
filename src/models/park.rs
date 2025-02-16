use crate::schema::parks;
use diesel::prelude::Queryable;
use diesel::prelude::Selectable;
use diesel::prelude::Insertable;
use diesel::prelude::AsChangeset;
use serde::{Deserialize, Serialize};
use uuid::Uuid;

#[derive(Queryable, Selectable, Insertable, AsChangeset, Serialize, Deserialize, Debug)]
#[diesel(table_name = parks)]
pub struct Park {
    pub park_id: Uuid,
    pub park_name: String,
    pub address: String,
    pub price_ph: f64,
    pub status: String,
    pub available_spots_count: i32,
    pub total_spots_count: i32,
    pub electro_charging_available: bool,
    pub rating: Option<f64>,
    pub park_balance: Option<f64>,
    pub latitude: f64,
    pub longitude: f64,
}
