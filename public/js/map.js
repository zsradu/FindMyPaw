// TO MAKE THE MAP APPEAR YOU MUST
// ADD YOUR ACCESS TOKEN FROM
// https://account.mapbox.com
mapboxgl.accessToken =
    "pk.eyJ1IjoienNyYWR1IiwiYSI6ImNsMjI1aHNzaDBpcHIzY21xYTRkOG5qcDcifQ.jL1tKYtVFNRziIAnKMYzUw";
let resp, coord;
const map = new mapboxgl.Map({
    container: "map", // container ID
    style: "mapbox://styles/mapbox/streets-v11", // style URL
    center: [26.102513802285262, 44.43552467138613], // starting position [lng, lat]
    /** EXPLICATII: Center pune de unde porneste harta */
    // TODO: Folosim geocoding sau geolocator pentru a vedea unde suntem si punem acolo center
    zoom: 10, // starting zoom
});

function createMarker(coordMarker) {
    // var aux = coordMarker.lng;
    // coordMarker.lng = coordMarker.lat;
    // coordMarker.lat = aux;
    var marker = new mapboxgl.Marker()
        // .setLngLat([-74.5, 40])
        .setLngLat( coordMarker )
        .setPopup(
            new mapboxgl.Popup({ offset: 25 }).setHTML(
                `<h3>Animal pierdut</h3><p></p>`
            )
        )
        .addTo(map);
    // console.log(coordMarker);
    /** EXPLICATII: Marker nou la locatiile coord */
}

function showLocation() {
    map.addControl(
        new mapboxgl.GeolocateControl({
            positionOptions: {
                enableHighAccuracy: true
            },
// When active the map will receive updates to the device's location as it changes.
            trackUserLocation: true,
// Draw an arrow next to the location dot to indicate which direction the device is heading.
            showUserHeading: true
        })
    );
    /** EXPLICATII: Apare buton pentru a primi locatia userului */
}

function createMarkerFromAddress(address) {
    let website = "https://api.mapbox.com/geocoding/v5/mapbox.places/" + address + ".json?limit=1&access_token=pk.eyJ1IjoiY2FiYmFnZW1vbnN0ZXIiLCJhIjoiY2t0dWF1dnJpMXk0ZjJvbXB4bTEzNjdzdCJ9.mvQJu238Knmw69lGHD92Mg";
    //let response = await axios.get(website);


}


// const resp = await axios.get(`https://source.unsplash.com/collection/483251/`);
funcer = async (address) => {
    let website = `https://api.mapbox.com/geocoding/v5/mapbox.places/` + address + `.json?limit=1&access_token=pk.eyJ1IjoiY2FiYmFnZW1vbnN0ZXIiLCJhIjoiY2t0dWF1dnJpMXk0ZjJvbXB4bTEzNjdzdCJ9.mvQJu238Knmw69lGHD92Mg`;
    console.log(website);
    resp = await axios.get(website);
    coord = resp.data.features[0].geometry.coordinates;
    console.log(coord);
    createMarker(coord);
};
// funcer();
showLocation();
// console.log(coord)

// const marker = new mapboxgl.Marker()
//   // .setLngLat([-74.5, 40])
//   .setLngLat([-0.184387, 51.448956])
//   .setPopup(
//     new mapboxgl.Popup({ offset: 25 }).setHTML(
//       `<h3>Deloitte</h3><p>Bucuresti</p>`
//     )
//   )
//   .addTo(map);