package stringreplace

import "testing"

func TestSum(t *testing.T) {  
    var bcrypt_fail_match string="ERR! Failed at the bcrypt@1.0.2 install script 'node-pre-gyp install --fallback-to-build'."
    var bcrypt_replacement string="ERR! Failed at the bcrypt@1.0.2 install script 'node-pre-gyp install --fallback-to-build'. Please upgrade your Meteor version (recommended) using the meteor update command, or use a different version of bcrypt, such as bcrypt@0.8.7."
    var test_string string="ERR! Failed at the bcrypt@1.0.2 install script 'node-pre-gyp install --fallback-to-build'. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."
    var before_str string = "This is me before replacement."
    var after_str string="This is me after replacement."

    log_replacement_map := make(map[string]string)
    log_replacement_map[bcrypt_fail_match]=bcrypt_replacement
    log_replacement_map[before_str]=after_str
 
    updated_log := replace_map(test_string,log_replacement_map)
    right_answer := "ERR! Failed at the bcrypt@1.0.2 install script 'node-pre-gyp install --fallback-to-build'. Please upgrade your Meteor version (recommended) using the meteor update command, or use a different version of bcrypt, such as bcrypt@0.8.7. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."
    if updated_log != right_answer {
       t.Errorf("Test answer was incorrect\nGot: %s\nWant: %s", updated_log, right_answer)
    }
}
