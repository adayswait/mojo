<template>
  <div id="stage">
    <canvas id="lifegame">Your browser doesn't support HTML5 element canvas.</canvas>
  </div>
</template>

<script>
/**
 * Life Game, a cellular automata
 * @author adayswait
 */
export default {
  data() {
    return {
      /**
       * @param life lives of each generation
       * @param lifeSize the size of each life shown on the screen
       * @param generationGap time(ms) between to generation
       * @param lifegame a canvas id
       */
      life: [],
      lifeSize: 5,
      generationGap: 100,
      world: undefined,
      worldSizex: 0,
      worldSizey: 0,
      stageWidth: 0,
      stageHeight: 0,
    };
  },
  methods: {
    /**
     * Create the first generation
     * @param life_ An array contains many lives
     * @param worldSizex_ How many lives are there in one row
     * @param worldSizey_ How many lives are there in one line
     */
    createLife(life_, worldSizex_, worldSizey_) {
      for (let worldx = 0; worldx < worldSizex_; worldx++) {
        life_[worldx] = [];
        for (let worldy = 0; worldy < worldSizey_; worldy++) {
          if (Math.random() > 0.7) {
            life_[worldx][worldy] = true;
          } else {
            life_[worldx][worldy] = false;
          }
        }
      }
      return life_;
    },

    /**
     * generate the next generation
     * @param life_ An array contains many lives
     * @param worldSizex_ How many lives are there in one row
     * @param worldSizey_ How many lives are there in one line
     */
    generateWorld(life_, worldSizex_, worldSizey_) {
      let nextLife = [];
      for (let worldx = 0; worldx < worldSizex_; worldx++) {
        nextLife[worldx] = [];
        for (let worldy = 0; worldy < worldSizey_; worldy++) {
          //calculate how many lives around current life
          (function (worldx_, worldy_) {
            let neighbour = 0;
            if (life_[worldx_ - 1]) {
              if (life_[worldx_ - 1][worldy_ - 1]) {
                neighbour += 1;
              }
              if (life_[worldx_ - 1][worldy_]) {
                neighbour += 1;
              }
              if (life_[worldx_ - 1][worldy_ + 1]) {
                neighbour += 1;
              }
            }
            if (life_[worldx_ + 1]) {
              if (life_[worldx_ + 1][worldy_ - 1]) {
                neighbour += 1;
              }
              if (life_[worldx_ + 1][worldy_]) {
                neighbour += 1;
              }
              if (life_[worldx_ + 1][worldy_ + 1]) {
                neighbour += 1;
              }
            }
            if (life_[worldx_][worldy_ - 1]) {
              neighbour += 1;
            }
            if (life_[worldx_][worldy_ + 1]) {
              neighbour += 1;
            }

            nextLife[worldx_][worldy_] = life_[worldx_][worldy_];
            //if there are 2 lives around current life
            //current life won't change its state
            if (neighbour == 2) {
              //nop
            }
            //if there are 3 lives around current life
            //current life will reborn(if current life is dead)
            else if (neighbour == 3) {
              nextLife[worldx_][worldy_] = true;
            }
            //otherwise, current life will die
            else {
              nextLife[worldx_][worldy_] = false;
            }
          })(worldx, worldy);
        }
      }
      return nextLife;
    },
    /**
     * show lives to world_
     * @param life_ An array contains many lives
     * @param worldSizex_ How many lives are there in one row
     * @param worldSizey_ How many lives are there in one line
     * @param world_ Canvas which created to show this game
     */
    visualWorld(life_, lifeSize_, worldSizex_, worldSizey_, world_) {
      world_.clearRect(0, 0, worldSizex_ * lifeSize_, worldSizey_ * lifeSize_);
      for (let worldx = 0; worldx < worldSizex_; worldx++) {
        for (let worldy = 0; worldy < worldSizey_; worldy++) {
          if (life_[worldx][worldy]) {
            world_.fillRect(
              lifeSize_ * worldx,
              lifeSize_ * worldy,
              lifeSize_ - 1,
              lifeSize_ - 1
            );
          }
        }
      }
    },
    timeFlies() {
      this.visualWorld(
        this.life,
        this.lifeSize,
        this.worldSizex,
        this.worldSizey,
        this.world
      );
      this.life = this.generateWorld(
        this.life,
        this.worldSizex,
        this.worldSizey
      );
    },
  },
  mounted() {
    this.stageWidth = document.getElementById("stage").clientWidth;
    this.stageHeight = document.getElementById("stage").clientHeight;
    window.console.log(this.stageWidth, this.stageHeight);
    document.getElementById("lifegame").setAttribute("width", this.stageWidth);
    document
      .getElementById("lifegame")
      .setAttribute("height", this.stageHeight);
    this.world = document.getElementById("lifegame").getContext("2d");
    this.worldSizex =
      Math.floor(document.getElementById("lifegame").offsetWidth / this.lifeSize);
    this.worldSizey =
      Math.floor(document.getElementById("lifegame").offsetHeight / this.lifeSize);
    this.world.fillStyle = "#000000";

    this.life = this.createLife(this.life, this.worldSizex, this.worldSizey);
    setInterval(this.timeFlies.bind(this), this.generationGap);
  },
};
</script>

<style scoped>
div {
  height: 100%;
}
</style>

--bthread_concurrency=6 --crash_on_fatal_log=true --raft_max_segment_size=8388608 --raft_sync=true --ip=10.1.1.43 --port=8602 --conf=10.1.1.248:8600:0,10.1.1.248:8601:0:0,10.1.1.43:8602:0,
