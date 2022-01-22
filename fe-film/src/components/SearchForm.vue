<template>
  <div class="input-group mb-3">
  <input type="text" class="form-control" placeholder="Search By Title" v-model="search">
  <button class="btn btn-secondary" type="button" @click=" search != '' ? searchFilm() : loadFilm()"><b-icon-search/> Cari Film</button>
</div>
</template>

<script>
export default {
data(){
    return{
        search:""
    }
},
methods:{
    searchFilm(){
        this.$store.commit('setSearch', this.search)
        this.$store.dispatch('loadSearchFilms', this.search)        
    },
    loadFilm(){
        this.$store.commit('setCurrentPage',1)
        this.$store.dispatch('loadFilms')
    }
}

}
</script>

<style>

</style>