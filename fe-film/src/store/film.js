import axios from "axios"
export const film = {
    state: () => ({
       films:[],
       totalPage:0,
       currentPage:1,
       search:'',
    }),

    actions: {
        loadFilms({commit, state}){
            axios.get('/api/films?page='+state.currentPage).then((response)=>{
                commit('setFilms', response.data.data)
                commit('setTotalPage', response.data.total_page)
                commit('setCurrentPage', response.data.current_page)
            })
        },
        loadSearchFilms({commit, state}, search){
            axios.get('/api/films/search/'+ search +'?page='+state.currentPage).then((response)=>{
                commit('setFilms', response.data.data)
                commit('setTotalPage', response.data.total_page)
                if (state.currentPage > response.data.total_page) {
                    commit('setCurrentPage', 1)
                }else{
                    commit('setCurrentPage', response.data.current_page)
                }
            })
        },
        
    },

    mutations: {
        setFilms(state, data){
            state.films = data
        },
        setTotalPage(state, data){
            state.totalPage = data
        },
        setCurrentPage(state, data){
            state.currentPage = data
        },
        setSearch(state, data){
            state.search = data
        }
    },

    getters: {
        getFilms( state ){
            return state.films
        },
        getTotalPage( state ){
            return state.totalPage
        },
        getCurrentPage( state ){
            return state.currentPage
        },
        getSearch(state){
            return state.search
        }
    }
}