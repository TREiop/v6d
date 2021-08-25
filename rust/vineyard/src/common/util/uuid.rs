use rand::Rng;

pub type ObjectID = u64;
pub type InstanceID = u64;
pub type Signature = u64;

// TODO: Rust parse check
pub fn object_id_from_string(s: &String) -> ObjectID {
    s.parse::<ObjectID>().unwrap()
}

pub fn object_id_to_string(id: ObjectID) -> String {
    String::new()
}

// pub fn get_blob_addr() {}

// pub fn generate_blob_id() {}


pub fn empty_blob_id() -> ObjectID {
    0x8000000000000000u64
}

pub fn generate_object_id() -> ObjectID {
    let mut rng = rand::thread_rng();
    0x7FFFFFFFFFFFFFFFu64 & rng.gen::<u64>()
}

pub fn generate_signature() -> Signature {
    let mut rng = rand::thread_rng();
    0x7FFFFFFFFFFFFFFFu64 & rng.gen::<u64>()
}

pub fn is_blob(id: ObjectID) -> bool {
    if let 0 = id & 0x8000000000000000u64 {
        return false;
    }
    true
}
