<template>
  <div>
    <v-toolbar floating dense style="width: 16em;  height: 40px;">
      <label>
        <gmap-autocomplete :value="location"
          @place_changed="setPlace">
        </gmap-autocomplete>
      </label>
    </v-toolbar>
    <gmap-map
      :center="center"
      :zoom="16"
      style="width:100%;  height: 50vh;"
    >
      <gmap-marker
        :key="index"
        v-for="(m, index) in markers"
        :position="m.position"
        @click="center=m.position"
      ></gmap-marker>
    </gmap-map>
  </div>
</template>

<script>
var comLoc
export default {
  name: "GoogleMap",
  data() {
    return {
      // default to Montreal to keep it simple
      // change this to whatever makes sense
      center: { lat: 22.346, lng: 114.186 },
      markers: [],
      places: [],
      currentPlace: null
    };
  },

  mounted() {
    this.geolocate();
  },
  props: {
    location: String
  },
  methods: {
    // receives a place object via the autocomplete component
    setPlace(place) {
      this.currentPlace = place;
      this.addMarker();
    },
    addMarker() {
      this.markers = []
      this.places = []
      if (this.currentPlace) {
        const marker = {
          lat: this.currentPlace.geometry.location.lat(),
          lng: this.currentPlace.geometry.location.lng()
        }
        this.markers.push({ position: marker })
        this.places.push(this.currentPlace)
        this.center = marker
        this.currentPlace = null
      }
    },
    geolocate: function() {
      navigator.geolocation.getCurrentPosition(position => {
        this.center = {
          lat: position.coords.latitude,
          lng: position.coords.longitude
        }
      })
    }
  }
};
</script>
