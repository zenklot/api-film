<template>
  <div>
    <button
      class="btn btn-success"
      data-bs-toggle="modal"
      data-bs-target="#staticBackdrop"
    >
      <b-icon-plus /> Add Film
    </button>

    <!-- Modal -->
    <div
      class="modal fade"
      id="staticBackdrop"
      data-bs-backdrop="static"
      data-bs-keyboard="false"
      tabindex="-1"
      aria-labelledby="staticBackdropLabel"
      aria-hidden="true"
    >
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="staticBackdropLabel">Tambah Film Baru</h5>
            <button
              type="button"
              class="btn-close"
              data-bs-dismiss="modal"
              aria-label="Close"
            ></button>
          </div>
          <div class="modal-body">
            <div class="input-group mb-3">
                <span class="input-group-text">Judul Film</span>
                <input type="text" class="form-control" placeholder="Judul Film" v-model="film.title">
            </div>
            <div class="input-group mb-3">
                <span class="input-group-text">Genre Film</span>
                <input type="text" class="form-control" placeholder="Drama, Aksi, Horor, ..." v-model="genre">
            </div>
            <div class="input-group mb-3">
                <span class="input-group-text">Rating Film</span>
                <input type="text" class="form-control" placeholder="1 - 10" v-model="film.rating">
            </div>
            <div class="input-group mb-3">
                <span class="input-group-text">Durasi Film</span>
                <input type="text" class="form-control" placeholder="70 min" v-model="film.duration">
            </div>
            <div class="input-group mb-3">
                <span class="input-group-text">Qualitas Film</span>
                <input type="text" class="form-control" placeholder="HD" v-model="film.quality">
            </div>
            <div class="input-group mb-3">
                <span class="input-group-text">Trailer Film Link</span>
                <input type="text" class="form-control" placeholder="https://www.youtube.com/watch?v=" v-model="film.trailer">
            </div>
            <div class="input-group mb-3">
                <span class="input-group-text">Streaming Film Link</span>
                <input type="text" class="form-control" placeholder="https://103.136.42.202/slug-judul-film" v-model="film.watch">
            </div>
            

          </div>
          <div class="modal-footer">
            <button
              type="button"
              class="btn btn-secondary"
              data-bs-dismiss="modal"
            >
              Tutup
            </button>
            <button type="button" class="btn btn-primary" @click="addFilm()">Tambah Film</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
    data(){
        return{
            genre:'',
            film:{
                title:'',
                genre:[],
                rating:'',
                duration:'',
                quality:'',
                trailer:'',
                watch:''
            }
        }
    },
    methods:{
        showAlert(icon, title, message = "") {
          this.$swal.fire(title, message, icon);
        },
        addFilm(){
             this.film.genre = this.genre.split(',')
            
            this.axios.post('api/film', this.film).then(()=>{
                this.showAlert('success', 'Tambah Film Baru Berhasil')
            }).catch(()=>{
                this.showAlert('error','Gagal Tambah Data', 'Mohon Periksa Lagi Setiap Inputan')
            })
        }
    }
};
</script>

<style>
</style>