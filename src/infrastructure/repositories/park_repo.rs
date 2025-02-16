use diesel::prelude::*;
use uuid::Uuid;
use crate::models::park::Park;
use crate::schema::parks::dsl::*;

pub struct ParkRepo;

impl ParkRepo {
    pub fn create(conn: &mut PgConnection, new_park: &Park) -> QueryResult<Park> {
        diesel::insert_into(parks)
            .values(new_park)
            .get_result(conn)
    }

    pub fn get_all(conn: &mut PgConnection) -> QueryResult<Vec<Park>> {
        parks.load::<Park>(conn)
    }

    pub fn get_by_id(conn: &mut PgConnection, park_id_val: Uuid) -> QueryResult<Park> {
        parks.filter(park_id.eq(park_id_val)).first(conn)
    }

    pub fn update(conn: &mut PgConnection, park_id_val: Uuid, park_data: &Park) -> QueryResult<Park> {
        diesel::update(parks.filter(park_id.eq(park_id_val)))
            .set(park_data)
            .get_result(conn)
    }

    pub fn delete(conn: &mut PgConnection, park_id_val: Uuid) -> QueryResult<usize> {
        diesel::delete(parks.filter(park_id.eq(park_id_val))).execute(conn)
    }
}
