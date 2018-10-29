var eventBus = new Vue()

const matchResult =  function(home,away) {
            if (home > away)
                return 1;
            else if (home === away)
                return 0;
            else
                return 2;
        }

var matchMixin = {
    methods: {
        pronoCorrect: (prono) => {
            return (matchResult(prono.home_score,prono.away_score) === matchResult(prono.prono_home_score, prono.prono_away_score))
        },
        pronoExact: (prono) => {
            return (prono.home_score == prono.prono_home_score && prono.away_score == prono.prono_away_score)
        }
    }
}

Vue.component('selector-tabs', {
    template: `
    <div>
        <div class="btn-group">
        <button  class="btn bg-secondary text-light" v-for="(tab, index) in tabs" 
          :key="index"
            @click="selectedTab = tab">
            {{ tab }}
            </button>
          </div>
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
    mixins: [matchMixin],
    template:`
    <div>
    <table class="table table-bordered table-striped">
    <thead>
    <tr>
        <th svope="col">#</th>
        <th scope="col">Home Team</th>
        <th scope="col">Prono Score</th>
        <th scope="col">Real Score</th>
        <th scope="col">Away Team</th>
        <th scope="col">Prono Correct</th>
        <th scope="col">Prono Exact</th>
    </tr>
    </thead>
    <tbody>
        <tr v-for="(prono,index) in pronos">
            <th scope="row">{{index}}</th>
                <td>{{ prono.home_team }}</td>
                <td>{{ prono.prono_home_score }} - {{ prono.prono_away_score }}</td>
                <td>{{ prono.home_score }} - {{ prono.away_score }}</td>
                <td>{{ prono.away_team }}</td>
                <td>
                    <div v-if=pronoCorrect(prono)>Yes</div>
                    <div v-else>No</div>
                </td>
                <td>
                    <div v-if=pronoExact(prono)>Yes</div>
                    <div v-else>No</div>
                </td>
            </tr>
    </tr>
    </tbody>
    </table>
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
    mixins: [matchMixin],
    template: `
    <div>
        General Stats 
        Tournament Stats:
        <ul class="list-group">
            <li class="list-group-item"> Home gains: {{ homeGains() }} </li>
            <li class="list-group-item"> Total Goal Scored: {{ goalsScoredTot() }} </li>
            <li class="list-group-item"> Average Goals Scored: {{ goalsScoredTot() / matchesTot }} </li>
        </ul>

        Prono Stats
        <ul class="list-group">
            <li class="list-group-item"> Correct Pronos : {{ matchesCorrect() }} / {{ matchesTot }} </li>
            <li class="list-group-item"> Home Gains predicted: {{ pronoHomeGains() }} </li>
            <li class="list-group-item"> Total Goals predicted: {{ goalsPronoTot() }} </li>
            <li class="list-group-item"> Average Goals Predicted: {{ goalsPronoTot() / matchesTot }} </li>
        </ul>
    <div>
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
                if (this.pronoCorrect(m)) {
                    return nb + 1
                } else { 
                    return nb
                }
            },0);
        },
        goalsScoredTot() {
            return this.pronos.reduce((nb,p) => nb + (p.home_score + p.away_score),0);
        },
        goalsPronoTot() {
            return this.pronos.reduce((nb,p) => nb + (p.prono_home_score + p.prono_away_score),0);
        },
        homeGains() {
            return this.pronos.filter(
                prono => matchResult(prono.home_score,prono.away_score) == 1).length;

        },
        pronoHomeGains() {
            return this.pronos.filter(
                prono => matchResult(prono.prono_home_score,prono.prono_away_score) == 1).length;

        }

    }
})

Vue.component('matches', {
    template:`
    <table class="table table-bordered table-striped">
    <thead>
    <tr>
        <th svope="col">#</th>
        <th scope="col">Home Team</th>
        <th scope="col">Home Score</th>
        <th scope="col">Away Score</th>
        <th scope="col">Away Team</th>
    </tr>
    </thead>
    <tbody>
        <tr v-for="(prono,index) in pronos">
            <th scope="row">{{index}}</th>
                <td>{{ prono.home_team }}</td>
                <td>{{ prono.home_score }}</td>
                <td>{{ prono.away_score }}</td>
                <td>{{ prono.away_team }}</td>
            </tr>
    </tr>
    </tbody>
    </table>
    `,
    props: {
        pronos: []
    }

})

Vue.component('tournaments',{
    template:`
    <div>
    Tournaments
        <ul class="list-group">
            <button class="list-group-item list-group-item-action" 
            v-for="(tourn,index) in tournaments"
                :key="index"
                @click="tournamentSelected(index)">
                {{ tourn.tournament }} {{ tourn.sport }}</button>
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
