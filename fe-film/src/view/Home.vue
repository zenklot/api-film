<template>
<div class="container">
  <h1>Halaman Home</h1>
  <div class="row mb-4">
      <div class="col-md-2" v-for="film in films" :key="film.id">
          <CardFilm :film="film" />
      </div>
  </div>

  <div class="row mt-5">
      <div class="col" v-for="index in totalPage" :key="index">
          <a href="" @click.prevent="nextPage(index)" class="badge bg-warning" v-if="currentPage != index">{{index}}</a>
          <span href="" class="badge bg-dark" v-else>{{index}}</span>
      </div>
  </div>
</div>
  
</template>

<script>
import CardFilm from '../components/CardFilm.vue'
export default {
    components:{
        CardFilm
    },
computed:{
    films(){
        return this.$store.getters.getFilms
    },
    currentPage(){
        return this.$store.getters.getCurrentPage
    },
    totalPage(){
        return this.$store.getters.getTotalPage
    }
},
mounted(){
    this.$store.dispatch('loadFilms', 1)
},
methods:{
    nextPage(page){
        this.$store.dispatch('loadFilms', page)
    }
}
}
</script>

<style>

</style>