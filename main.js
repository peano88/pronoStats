

var eventBus = new Vue()

Vue.component('selector-tabs', {
    template: `
    <div>
        <ul>
          <span class="tab" 
          v-for="(tab, index) in tabs" 
          :key="index"
          @click="selectedTab = tab">{{ tab }}
          </span>
          </ul>

        <div v-show="selectedTab === 'Matches'">
            <matches :pronos=pronos></matches>            
        </div>
        <div v-show="selectedTab === 'Pronos'">
            <pronos :pronos=pronos></pronos>            
        </div>
        <div v-show="selectedTab === 'Stats'">
            <stats :pronos=pronos></stats>            
        </div>
    </div>
      `,
    props: {
        pronos: []
    },
    data() {
        return {
            tabs: ['Matches', 'Pronos', 'Stats'],
            selectedTab: 'Matches'
        }
    }
})

Vue.component('newProno', {
    template:`
    <div>
        <form @submit.prevent="onSubmit">
            <p>
                <label for="pronoHomeScore">Home Score</label>
                <input id="pronoHomeScore" v-model.number="prono_home_score" type="number">
            </p>
            <p>
                <label for="pronoAwayScore">Away Score</label>
                <input id="pronoAwayScore" v-model.number="prono_away_score" type="number">
            </p>
            <p>
                <input type="submit" value="Submit">
            </p>
        </form>
    </div>
    `,
    data() {
        return {
            home_team: "FixedA",
            away_team: "FixedB",
            home_score: 3,
            away_score: 1,
            prono_home_score: 0,
            prono_away_score: 0
        }
    },
    methods: {
        onSubmit() {
            let thisProno = {
                home_team: this.home_team,
                away_team: this.away_team,
                home_score: this.home_score,
                away_score: this.away_score,
                prono_home_score: this.prono_home_score,
                prono_away_score: this.prono_away_score
            }
                this.home_team = "FixedA"
                this.away_team = "FixedB"
                this.home_score = 3
                this.away_score = 1
                this.prono_home_score = 0
                this.prono_away_score = 0

            this.$emit('new-prono',thisProno)
            
        }

    }
        
})

Vue.component('pronos', {
    template:`
    <div>
        <div v-for="prono in pronos">
            {{ prono.home_team }} - {{ prono.away_team }}</br>
            {{ prono.prono_home_score }} : {{ prono.prono_away_score }}
            {{ prono.home_score }} : {{ prono.away_score }}
        </div>
        <button v-on:click="formVisible">New</button>
        <div v-show="formV">
            <newProno @new-prono="newProno"></newProno>
        </div>
    </div>
    `,
    props: {
        pronos: []
    },
    data() {
        return {
            formV: false
        }
    },
    methods: {
        formVisible() {
            this.formV = true;
        },
        newProno(aProno) {
            this.formV = false
            eventBus.$emit('new-prono', aProno)
        }

    }

})

Vue.component('stats', {
    template: `
    <div>
        You got {{ matchesCorrect() }} matches out of {{ matchesTot }} <br>
        Home team won {{ homeGains() }} and you predicted {{ pronoHomeGains() }} times <br>

        Total goals scored : {{ goalsScoredTot() }}  <br>
        Total goals predicted : {{ goalsPronoTot() }}  <br>
        Average goals: {{ goalsScoredTot() / matchesTot }} <br>
        Average goals predicted: {{ goalsPronoTot() / matchesTot }}
    </div>
    `,
    props: {
        pronos: []
    },
    computed: {
        matchesTot: function () {
            return this.pronos.length;
        }
    },
    methods: {
        matchesCorrect: function () {
            return this.pronos.reduce((nb,m) => {
                if (this.matchResult(m.home_score,m.away_score) === 
                    this.matchResult(m.prono_home_score,m.prono_away_score)) {
                    return nb + 1
                } else { 
                    return nb
                }
            },0);
        },
        matchResult:  (home,away) => {
            if (home > away)
                return 1;
            else if (home === away)
                return 0;
            else
                return 2;
        },
        goalsScoredTot() {
            return this.pronos.reduce((nb,p) => nb + (p.home_score + p.away_score),0);
        },
        goalsPronoTot() {
            return this.pronos.reduce((nb,p) => nb + (p.prono_home_score + p.prono_away_score),0);
        },
        homeGains() {
            return this.pronos.filter(
                prono => this.matchResult(prono.home_score,prono.away_score) == 1).length;

        },
        pronoHomeGains() {
            return this.pronos.filter(
                prono => this.matchResult(prono.prono_home_score,prono.prono_away_score) == 1).length;

        }

    }
})

Vue.component('matches', {
    template:`
    <div>
        <div v-for="prono in pronos">
            {{ prono.home_team }} - {{ prono.away_team }}</br>
            {{ prono.home_score }} : {{ prono.away_score }}
        </div>
    </div>
    `,
    props: {
        pronos: []
    }

})

Vue.component('tournaments',{
    template:`
    <div>
        <ul>
            <li class="tab" v-for="(tourn,index) in tournaments"
                :key="index"
                @click="tournamentSelected(index)" >
                <span>{{ tourn.tournament }} {{ tourn.sport }}</span>
            </li>
        </ul>
    </div>
    `,
    data() {
        return {
            activeTournament: ""
        }
    },
    props: {
        tournaments: []
    },
    methods: {
        tournamentSelected(i) {
            activeTournament = this.tournaments[i].tournament
            eventBus.$emit('sel-tourn',i)
        }
    }
})

var app = new Vue({
    el: '#app',
    data: {
        tournamentName: '',
        tournaments: [],
        pronos: []
    },
    mounted() {
        this.getPronos()
        eventBus.$on('new-prono', prono => {
            let url = "http://localhost:4000/tournaments/5b8a42af5a9f6508e6e5efb7/prono"
            axios.post(url,prono)
                .then(function (response) {
                    console.log(response);
                })
                .catch(function (error) {
                    console.log(error);
                });
        })
        eventBus.$on('sel-tourn', index => {
            this.pronos = this.tournaments[index].pronos
            this.tournamentName = this.tournaments[index].tournament
        })

    },
    methods: {
        getPronos() {
            //let url = "http://localhost:4000/tournaments/5b8a42af5a9f6508e6e5efb7"
            let url = "http://localhost:4000/tournaments/?user=UserTest"
            axios.get(url).then((response) => {
                this.tournaments = response.data
                this.pronos = response.data[0].pronos;
                this.tournamentName = response.data[0].tournament
            }).catch( error => { console.log(error) })
        }
    }
})
