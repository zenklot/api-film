<template>
  <div class="card mb-2 ms-2">
    <a :href="film.trailer" target="_blank">
      <img :src="getYtImg(film.trailer)" class="card-img-top" alt="" />
    </a>
    <div class="card-body">
      <h5 class="card-title">{{ film.title }}</h5>
      <div class="row text-start">
        <div class="col-12">
          <div
            class="badge bg-secondary ms-1"
            v-for="(genre, index) in film.genre"
            :key="index"
          >
            {{ genre }}
          </div>
        </div>
        <div class="col-12">
          <b-icon-star-fill /> Rating : {{ film.rating }}
        </div>
        <div class="col-12">
          <b-icon-alarm /> Duration : {{ film.duration }}
        </div>
        <div class="col-12">
          <b-icon-camera-reels-fill /> Quality : {{ film.quality }}
        </div>
      </div>
    </div>
    <div class="card-footer d-flex bd-highlight">
      <a
        :href="film.watch"
        target="_blank"
        class="btn btn-primary me-auto bd-highlight"
        ><b-icon-play-circle-fill /> Streaming</a
      >
      <button class="btn btn-danger" @click="deleteFilmById(film.id)">
        <b-icon-trash-fill />
      </button>
    </div>
  </div>
</template>

<script>
export default {
  props: ["film"],
  methods: {
    showToast(icon, message) {
      const Toast = this.$swal.mixin({
        toast: true,
        position: "top-end",
        showConfirmButton: false,
        timer: 3000,
        timerProgressBar: true,
        didOpen: (toast) => {
          toast.addEventListener("mouseenter", this.$swal.stopTimer);
          toast.addEventListener("mouseleave", this.$swal.resumeTimer);
        },
      });
      Toast.fire({
        icon: icon,
        title: message,
      });
    },
    getYtImg(link) {
      return (
        "https://img.youtube.com/vi/" +
        link.replace("https://www.youtube.com/watch?v=", "") +
        "/0.jpg"
      );
    },
    deleteFilmById(id) {
      this.$swal
        .fire({
          title: "Yakin Menghapus Film Ini?",
          text: "Kamu tidak dapat mengembalikan film in setelah dihapus!",
          icon: "warning",
          showCancelButton: true,
          confirmButtonColor: "#3085d6",
          cancelButtonColor: "#d33",
          confirmButtonText: "Yes, delete it!",
        })
        .then((result) => {
          if (result.isConfirmed) {
            this.axios
              .delete("/api/film/" + id)
              .then(() => {
                this.showToast("success", "Film Deleted");
                this.$store.dispatch("loadFilms");
              })
              .catch(() => {
                this.showToast("error", "Error Delete Film");
              });
          }
        });
    },
  },
};
</script>
