
const domain = window.location.hostname;
const port = 8700;

export const Post = (endpoint: string, body: string): Promise<Response> => {
    let url = `http://${domain}:${port}/api/${endpoint}`;
    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: body
    };
    return fetch(url, requestOptions)
        .then(response => {
            if (response.status > 204) {
                throw new Error('Not authorized')
            }
            return response.json()
        });
}

export const Get = (endpoint: string): Promise<Object> => {
    let url = `http://${domain}:${port}/api/${endpoint}`;
    const requestOptions = {
        method: 'GET',
        headers: { 'Content-Type': 'application/json' }
    };
    return fetch(url, requestOptions)
        .then(response => {
            if (response.status > 204) {
                throw new Error('Not authorized')
            }
            return response.json()
        });
}
