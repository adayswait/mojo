<template>
  <canvas
    id="lifegame"
    width="100%"
    height="100%"
  >Your browser doesn't support HTML5 element canvas.</canvas>
</template>

<script>
export default {
  methods: {},
  mounted() {
    /**
     * Life Game, a cellular automata
     * @author adayswait
     */

    /**
     * @param life lives of each generation
     * @param lifeSize the size of each life shown on the screen
     * @param generationGap time(ms) between to generation
     * @param lifegame a canvas id
     */
    var life = [];
    var lifeSize = 2;
    var generationGap = 100;
    var world = document.getElementById("lifegame").getContext("2d");
    var worldSizex = world.width / lifeSize;
    var worldSizey = world.height / lifeSize;
    world.fillStyle = "#222222";

    life = createLife(life, worldSizex, worldSizey);
    setInterval(time, generationGap);

    function time() {
      visualWorld(life, lifeSize, worldSizex, worldSizey, world);
      life = generateWorld(life, worldSizex, worldSizey);
    }

    /**
     * Create the first generation
     * @param life_ An array contains many lives
     * @param worldSizex_ How many lives are there in one row
     * @param worldSizey_ How many lives are there in one line
     */
    function createLife(life_, worldSizex_, worldSizey_) {
      for (var worldx = 0; worldx < worldSizex; worldx++) {
        life_[worldx] = [];
        for (var worldy = 0; worldy < worldSizey_; worldy++) {
          if (Math.random() > 0.7) {
            life_[worldx][worldy] = true;
          } else {
            life_[worldx][worldy] = false;
          }
        }
      }
      return life_;
    }

    /**
     * generate the next generation
     * @param life_ An array contains many lives
     * @param worldSizex_ How many lives are there in one row
     * @param worldSizey_ How many lives are there in one line
     */
    function generateWorld(life_, worldSizex_, worldSizey_) {
      var nextLife = [];
      for (var worldx = 0; worldx < worldSizex_; worldx++) {
        nextLife[worldx] = [];
        for (var worldy = 0; worldy < worldSizey_; worldy++) {
          //calculate how many lives around current life
          (function (worldx_, worldy_) {
            var neighbour = 0;
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
    }
    /**
     * show lives to world_
     * @param life_ An array contains many lives
     * @param worldSizex_ How many lives are there in one row
     * @param worldSizey_ How many lives are there in one line
     * @param world_ Canvas which created to show this game
     */
    function visualWorld(life_, lifeSize_, worldSizex_, worldSizey_, world_) {
      world_.clearRect(0, 0, worldSizex_ * lifeSize_, worldSizey_ * lifeSize_);
      for (var worldx = 0; worldx < worldSizex_; worldx++) {
        for (var worldy = 0; worldy < worldSizey_; worldy++) {
          if (life_[worldx][worldy]) {
            world_.fillRect(
              lifeSize_ * worldx,
              lifeSize_ * worldy,
              lifeSize_,
              lifeSize_
            );
          }
        }
      }
    }
  },
};
</script>

<style scoped>
div {
  font-size: xx-large;
}
</style>

