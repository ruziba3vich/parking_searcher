// @generated automatically by Diesel CLI.

diesel::table! {
    history (history_id) {
        history_id -> Uuid,
        user_id -> Uuid,
        park_id -> Uuid,
        spot_id -> Uuid,
        entry_time -> Timestamp,
        exit_time -> Nullable<Timestamp>,
        total_cost -> Nullable<Float8>,
        #[max_length = 20]
        status -> Varchar,
    }
}

diesel::table! {
    parks (park_id) {
        park_id -> Uuid,
        #[max_length = 64]
        park_name -> Varchar,
        #[max_length = 166]
        address -> Varchar,
        price_ph -> Float8,
        #[max_length = 20]
        status -> Varchar,
        available_spots_count -> Int4,
        total_spots_count -> Int4,
        electro_charging_available -> Bool,
        rating -> Nullable<Float8>,
        park_balance -> Nullable<Float8>,
        latitude -> Float8,
        longitude -> Float8,
    }
}

diesel::allow_tables_to_appear_in_same_query!(
    history,
    parks,
);
