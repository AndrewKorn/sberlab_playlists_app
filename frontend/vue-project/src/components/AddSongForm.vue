<template>
  <form class="form-inline" @submit.prevent="onSubmit">
    <div class="form-group mx-sm-3 mb-2">
      <input type="text" class="form-control" id="inputArtist" placeholder="Artist" v-model="artist">
    </div>
    <div class="form-group mx-sm-3 mb-2">
      <input type="text" class="form-control" id="inputName" placeholder="Song name" v-model="song_name">
    </div>
    <button class="btn btn-secondary mb-2">Add song</button>
  </form>
</template>

<script>
export default {
  name: "AddSongForm",
  data() {
    return {
      sound_cloud_id: "",
      artist: "",
      song_name: "",
      duration: ""
    }
  },

  methods: {
    onSubmit() {
      if (this.song_name.trim() && this.artist.trim()) {
        const newSong = {
          sound_cloud_id: this.sound_cloud_id,
          artist: this.artist,
          song_name: this.song_name,
          duration: this.duration
        }
        this.$emit("createSong", newSong)
      }
      this.song_name = ''
      this.artist = ''
    }
  }
}
</script>

<style scoped>
form {
  margin: 0 auto;
  width: 600px;
}
</style>