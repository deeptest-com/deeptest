const axios = require('axios')

export async function uploadFile(url, file, params) {
    const data = new FormData();
    data.append('file', file);

    Object.keys(params).forEach(key=>{
        data.append(key, params[key]);
    })

    const config = {
        onUploadProgress: function(progressEvent) {
            const percentCompleted = Math.round( (progressEvent.loaded * 100) / progressEvent.total );
        }
    };

    const res = await axios.post(url, data, config)
    console.log('===', res)
}
