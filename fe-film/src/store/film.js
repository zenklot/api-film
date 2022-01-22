import axios from "axios"
export const film = {
    state: () => ({
       films:[],
       totalPage:0,
       currentPage:1
    }),

    actions: {
        loadFilms({commit}, page){
            axios.get('/api/films?page='+page).then((response)=>{
                commit('setFilms', response.data.data)
                commit('setTotalPage', response.data.total_page)
                commit('setCurrentPage', response.data.current_page)
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
        }
    }
}