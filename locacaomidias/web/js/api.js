async function GET(url) {
    const res = await fetch(url);
    return res.json();
}
