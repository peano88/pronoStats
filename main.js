var app = new Vue({
    el: '#app',
    data: {
        tournamentName: '',
        pronos: [
            /*{
                tournament : "Mickey's tournament",
                home_team : "Goofy",
                away_team : "Donald Duck",
                home_score : "1",
                away_score : "2",
                prono_home_score : "1",
                prono_away_score : "1"
            },
            {
                tournament : "Mickey's tournament",
                home_team : "Uncle Scrooge",
                away_team : "Rockerduck",
                home_score : "1",
                away_score : "0",
                prono_home_score : "2",
                prono_away_score : "1"
            }*/
        ]
    },
    mounted: function () {
        this.getPronos()
    },
    methods: {
        getPronos() {
            let url = "http://localhost:4000/tournament/5b64ef739d5a1f22c9e7d60b"
            axios.get(url).then((response) => {
                this.pronos = response.data.pronos;
                this.tournamentName = response.data.tournament
            }).catch( error => { console.log(error) })
        }
    }
})
