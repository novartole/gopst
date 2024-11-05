use std::ffi::CStr;

const PARSING_FAILED: u8 = 100;
const INVALID_UTF8: u8 = 101;

#[no_mangle]
#[allow(clippy::missing_safety_doc)]
pub unsafe extern "C" fn launch(input: *const libc::c_char) -> libc::c_uchar {
    match CStr::from_ptr(input).to_str().map(str::to_string) {
        Ok(s) => match shlex::split(&s) {
            Some(args) => {
                let output = typstlib::exec(args);
                *(&output as *const _ as *const u8)
            }
            None => PARSING_FAILED,
        },
        Err(why) => {
            println!("{}", why);
            INVALID_UTF8
        }
    }
}
