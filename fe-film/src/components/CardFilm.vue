<template>
    <div class="card">
<a :href="film.trailer" target="_blank">
  <img :src="getYtImg(film.trailer)" class="card-img-top" alt="">
  </a>
  <div class="card-body">
    <h5 class="card-title">{{ film.title }}</h5>
    <div class="row text-start">
        <div class="col-12">
            <div class="badge bg-secondary ms-1" v-for="(genre,index) in film.genre" :key="index">{{genre}}</div>
        </div>
        <div class="col-12">
            <b-icon-star-fill /> Rating : {{film.rating}}
        </div>
        <div class="col-12">
            <b-icon-alarm /> Duration : {{film.duration}}
        </div>
        <div class="col-12">
            <b-icon-camera-reels-fill /> Quality : {{film.quality}}
        </div>
    </div>
    
    <a :href="film.watch" target="_blank" class="btn btn-primary"><b-icon-play-circle-fill /> Streaming</a>
  </div>
</div>
</template>

<script>
export default {
 props:['film'],
 methods:{
     getYtImg(link){
         return 'https://img.youtube.com/vi/' + link.replace('https://www.youtube.com/watch?v=','') + '/0.jpg'
     }
 }
}
</script>
