<script setup lang="ts">
import { ref, watch, onMounted } from 'vue'
import type { Node, Edge } from '@vue-flow/core'  
import { VueFlow, useVueFlow } from '@vue-flow/core'


import SpecialNode from './components/SpecialNode.vue'
import SpecialEdge from './components/SpecialEdge.vue'
import Background from './components/Background.vue'
import MenuBar from './components/MenuBar.vue'
// these are our nodes
// const nodes = ref<Node[]>([
//   // an input node, specified by using `type: 'input'`
//   {
//     id: '1',
//     type: 'input',
//     position: { x: 250, y: 5 },
//     // all nodes can have a data object containing any data you want to pass to the node
//     // a label can property can be used for default nodes
//     data: { label: 'Node 1' },
//   },

//   // default node, you can omit `type: 'default'` as it's the fallback type
//   {
//     id: '2',
//     position: { x: 100, y: 100 },
//     data: { label: 'Node 2' },
//   },

//   // An output node, specified by using `type: 'output'`
//   {
//     id: '3',
//     type: 'output',
//     position: { x: 400, y: 200 },
//     data: { label: 'Node 3' },
//   },

//   // this is a custom node
//   // we set it by using a custom type name we choose, in this example `special`
//   // the name can be freely chosen, there are no restrictions as long as it's a string
//   // {
//   //   id: '4',
//   //   type: 'special', // <-- this is the custom node type name
//   //   position: { x: 400, y: 200 },
//   //   data: {
//   //     label: 'Node 4',
//   //     hello: 'world',
//   //   },
//   // },
// ])

const nodes = ref<Node[]>([])
const edges = ref<Edge[]>([])
// these are our edges
// const edges = ref<Edge[]>([
//   // default bezier edge
//   // consists of an edge id, source node id and target node id
//   {
//     id: 'e1->2',
//     source: '1',
//     target: '2',
//   },

//   // set `animated: true` to create an animated edge path
//   {
//     id: 'e2->3',
//     source: '2',
//     target: '3',
//     animated: true,
//   },

//   // a custom edge, specified by using a custom type name
//   // we choose `type: 'special'` for this example
//   // {
//   //   id: 'e3->4',
//   //   type: 'special',
//   //   source: '3',
//   //   target: '4',

//   //   // all edges can have a data object containing any data you want to pass to the edge
//   //   data: {
//   //     hello: 'world',
//   //   }
//   // },
// ])

const ready = ref(false)

const instance = useVueFlow()
const { getNodes, onPaneReady, getEdges } = instance

// watch the stored nodes
watch(getNodes, (nodes) => console.log('nodes changed', nodes))
watch(getEdges, (edges) => console.log('edges changed', edges))

const addNode  = async () => {
  const id = (nodes.value.length + 1).toString()
  const position = {
    x: 0,
    y: 100,
  }

  const node = {
    id,
    type: 'default',
    position,
    data: { label: `Node ${id}` },
  }
  nodes.value.push(node)

  const goNode = {
    id: node.id,
    type: node.type,
    x: node.position.x,
    y: node.position.y,
    label: node.data.label,
  }

  // send post request to server to add node
  const response = await fetch("http://localhost:3000/add-node", {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(goNode),
  });
  console.log("response: ", response)
  const data = await response.json();
  console.log('Data fetched:', data);
}

// function addNode = () {
//   const id = (nodes.value.length + 1).toString()
//   const position = {
//     x: 0,
//     y: 100,
//   }
  
//   nodes.value.push({
//     id,
//     type: 'default',
//     position,
//     data: { label: `Node ${id}` },
//   })

// }
function removeEdge(id) {
  edges.value = edges.value.filter((edge) => edge.id !== id)
}

async function loadFlowState() {
  const response = await fetch("http://localhost:3000/load-flow-state", {
    method: 'GET',
  });
  const data = await response.json();
  console.log('Data fetched:', data);
  for (const node of data.nodes) {
    nodes.value.push({
      id: node.Id,
      type: node.Type,
      position: {
        x: node.X,
        y: node.Y,
      },
      data: {
        label: node.Label,
      },
    })
  }
  for (const edge of data.edges) {
    edges.value.push({
      id: edge.Id,
      source: edge.Source,
      target: edge.Target,
      animated: edge.Animated,
    })
  }
}

onMounted(async () => {
  await loadFlowState()
  ready.value = true
})

console.log("nodes: ", nodes.value)
console.log("edges: ", edges.value)

onPaneReady((i) => i.fitView())
</script>

<template>
  <Background />
  <p style="color: white; font-size: 60px; text-align: center; margin-top: 200px;" v-if="!ready">Loading...</p>
  <div class="base-flow-container" v-if="ready">
    <VueFlow :nodes="nodes" :edges="edges">
      
    </VueFlow>
    <MenuBar @add-node="addNode" />
  </div>
</template>

<style>
/* import the necessary styles for Vue Flow to work */
@import '@vue-flow/core/dist/style.css';

/* import the default theme, this is optional but generally recommended */
@import '@vue-flow/core/dist/theme-default.css';


.base-flow-container {
  width: 100vw;
  height: 100vh;
  display: flex;
  place-items: center;
  padding: 0.2rem;
  box-sizing: border-box;  /* Add this line */
}

.vue-flow__node-failure {  /* Changed from custom to special */
  background: red;
  color: white;
  border: 1px solid red;
  border-radius: 4px;
  box-shadow: 0 0 0 1px red;
  padding: 8px;
}

.vue-flow__node-success {  /* Changed from custom to special */
  background: green;
  color: white;
  border: 1px solid black;
  border-radius: 4px;
  box-shadow: 0 0 0 1px black;
  padding: 8px;
}

</style>


      <!-- bind your custom node type to a component by using slots, slot names are always `node-<type>`
        <template #node-special="specialNodeProps">
          <SpecialNode v-bind="specialNodeProps" />
        </template>
  
        <!-- bind your custom edge type to a component by using slots, slot names are always `edge-<type>` -->
        <!-- <template #edge-special="specialEdgeProps">
          <SpecialEdge v-bind="specialEdgeProps" />
        </template> --> -->