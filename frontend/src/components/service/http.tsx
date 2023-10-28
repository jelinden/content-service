
const protocol = window.location.protocol;
const domain = window.location.hostname;
const port = window.location.port;

export const Post = (endpoint: string, body: string): Promise<Response> => {
    let url = `${protocol}://${domain}:${port}/api/${endpoint}`;
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
    let url = `${protocol}://${domain}:${port}/api/${endpoint}`;
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

export const Delete = (endpoint: string, id: number): Promise<Object> => {
    let url = `${protocol}://${domain}:${port}/api/${endpoint}/${id}`;
    const requestOptions = {
        method: 'DELETE',
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