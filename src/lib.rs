extern crate wasm_bindgen;
extern crate markdown;

use wasm_bindgen::prelude::*;

#[wasm_bindgen]
pub fn parse(input: &str) -> String {
    let result = markdown::to_html(input);
    result
}
