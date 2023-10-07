
class ValidationFields {
    constructor(username: string, password: string) {
        this.username = username
        this.password = password
    }
    username: string
    password: string
}

const validateFields = (username: string, password: string): ValidationFields => {
    var validations = new ValidationFields('', '')
    if (username === undefined || username.length < 5 || username.length > 127) {
        validations.username = 'Minimum 5, maximum 128 characters';
    } 
    if (password === undefined || password.length < 5 || password.length > 127) {
        validations.password = 'Minimum 5, maximum 128 characters';
    }
    return validations
}

export default validateFields