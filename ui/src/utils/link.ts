
export const openHelp = (page, bookmark?) => {
    const url = process.env.VUE_HELP_URL;
    window.open(`${url}/${page}.html`);
}
