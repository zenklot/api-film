<template>
<div class="container">
  <h1>REST API LIST FILM</h1>
  <div class="row mt-5 justify-content-center">
      <div class="col-md-6 col-sm-10">
          <SearchForm />
      </div>
  </div>
  <div class="d-flex justify-content-end">
        <AddFilm />
  </div>
  <div class="row justify-content-center row-cols-1 row-cols-md-5 row-col-sm-4 mb-4 mt-3" v-if="films.length > 0">
        <CardFilm v-for="film in films" :key="film.id" :film="film" /> 
  </div>
  <div class="row" v-else>
      <div class="col">
          <p class="fs-3 fw-bold">Opps. Belum Ada Film di Database!</p>
      </div>
  </div>

  <nav>
  <ul class="pagination">
    <li :class="currentPage != index ? 'page-item ' : 'page-item active'" v-for="index in totalPage" :key="index"><a class="page-link" href="" @click.prevent="nextPage(index)">{{index}}</a></li>
  </ul>
</nav>

<footer class="footer mt-auto py-3 bg-light">
  <div class="container">
    <span class="text-muted">2022 By Raisa Supriatna</span>
  </div>
</footer>
</div>
  
</template>

<script>
import CardFilm from '../components/CardFilm.vue'
import SearchForm from '../components/SearchForm.vue'
import AddFilm from '../components/AddFilm.vue'
export default {
    components:{
        CardFilm,
        SearchForm,
        AddFilm
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
    },
    search(){
        return this.$store.getters.getSearch
    }
},
mounted(){
    this.$store.dispatch('loadFilms', 1)
},
methods:{
    nextPage(page){
        this.$store.commit('setCurrentPage', page)
        if (this.search != '') {
            this.$store.dispatch('loadSearchFilms', this.search, page)
        }else{
            this.$store.dispatch('loadFilms')
        }
    }

}
}
</script>

<style>

</style>