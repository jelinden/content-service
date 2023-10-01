
const domain = window.location.hostname;
const port = 8700;

const Post = (endpoint: string, body: string): Promise<Response> => {
    let url = `http://${domain}:${port}/api/${endpoint}`;
    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: body
    };
    return fetch(url, requestOptions)
        .then(response => response.json());
}

export default Post