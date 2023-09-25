
export const openHelp = (page, bookmark?) => {
    const base = process.env.VUE_APP_HELP_URL;
    const url = `${base}/${page}.html`
    console.log('process.env', process.env)
    window.open(url);
}
