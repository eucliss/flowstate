<script setup lang="ts">
import { ref, watch, onMounted } from 'vue'
import type { Node, Edge } from '@vue-flow/core'  
import { VueFlow, useVueFlow } from '@vue-flow/core'

import SpecialNode from './components/SpecialNode.vue'
import SpecialEdge from './components/SpecialEdge.vue'
import Background from './components/Background.vue'
import MenuBar from './components/MenuBar.vue'

const nodes = ref<Node[]>([])
const edges = ref<Edge[]>([])
const ready = ref(false)

const instance = useVueFlow()
const { 
  getNodes, 
  onPaneReady, 
  getEdges, 
  getSelectedNodes, 
  getConnectedEdges,
  getIncomers,
  onEdgeClick, 
  onEdgesChange,
  onNodeDragStop,
  onConnect
} = instance

// onEdgesChange((event) => {
//   console.log("edges changed: ", event)
// })

onNodeDragStop(async (event) => { 
  console.log("node dragged: ", event)
  const node = {
    id: event.node.id,
    x: Math.floor(event.node.position.x),
    y: Math.floor(event.node.position.y),
    label: event.node.data.label,
    type: event.node.type,
  }
  const response = await fetch("http://localhost:3000/update-node", {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(node),
  });
  console.log("response: ", response)
  const data = await response.json();
  console.log('Data fetched:', data);
  await loadFlowState()
})

onConnect(async (event) => {
  const new_edge = {
    id: event.source + "->" + event.target,
    source: event.source,
    target: event.target,
    animated: true,
  }
  edges.value.push(new_edge)
  // Connected a new edge
  const response = await fetch("http://localhost:3000/connect-edge", {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(new_edge),
  });
  console.log("response: ", response)
  const data = await response.json();
  console.log('Data fetched:', data);
  console.log("connect: ", event)
  await loadFlowState()
})

// watch the stored nodes
// watch(getNodes, (nodes) => console.log('nodes changed', nodes))
// watch(getEdges, (edges) => console.log('edges changed watching', edges))

const addNode  = async (nodeData: { name: string }) => {
  console.log("adding a new node, len of nodes: ", nodes.value.length)
  // const id = (nodes.value.length + 1).toString()
  const id = crypto.randomUUID()
  const position = {
    x: 0,
    y: 100,
  }

  const node = {
    id,
    type: 'default',
    position,
    data: { label: nodeData.name },
  }
  nodes.value.push(node)

  const goNode = {
    id: node.id,
    type: node.type,
    x: node.position.x,
    y: node.position.y,
    label: nodeData.name,
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
  await loadFlowState()
}

function removeEdge(id) {
  edges.value = edges.value.filter((edge) => edge.id !== id)
}

async function loadFlowState() {
  const response = await fetch("http://localhost:3000/load-flow-state", {
    method: 'GET',
  });
  const data = await response.json();
  console.log('Data fetched:', data);
  const temp_nodes = []
  for (const node of data.nodes) {
    temp_nodes.push({
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
  const temp_edges = []
  for (const edge of data.edges) {
    temp_edges.push({
      id: edge.Id,
      source: edge.Source,
      target: edge.Target,
      animated: edge.Animated,
    })
  }
  nodes.value = temp_nodes
  edges.value = temp_edges
}

onMounted(async () => {
  await loadFlowState()
  ready.value = true
  window.addEventListener('keydown', handleKeyDown)
})

const handleKeyDown = async (event) => {
  console.log("selected nodes: ", getSelectedNodes.value)
  if (getSelectedNodes.value.length == 0) {
    return
  }
  const selected_node = getSelectedNodes.value[0].id
  const connections = getConnectedEdges(selected_node)
  const connection_ids = connections.map((edge) => edge.id)
  if (event.keyCode === 46 || event.keyCode === 8) {
    var response = await fetch("http://localhost:3000/delete-node", {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(selected_node),
  });
  console.log("response: ", response)
  var data = await response.json();
  console.log('Data fetched:', data);

  response = await fetch("http://localhost:3000/delete-edges", {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(connection_ids),
  });
  console.log("delete edges response: ", response)
  data = await response.json();
  console.log('delete edges data:', data);
  }
  await loadFlowState()

}

const resetFlowState = async () => {
  const response = await fetch("http://localhost:3000/reset", {
    method: 'GET',
  });
  const data = await response.json();
  console.log('Data fetched:', data);
  await loadFlowState()
}

console.log("nodes: ", nodes.value)
console.log("edges: ", edges.value)

onEdgeClick((event) => {
  console.log("edge clicked: ", event)
})

onPaneReady((i) => i.fitView())

const handleAddNode = (nodeData: { name: string }) => {
  console.log('App - Received node data:', nodeData)
  // Add your node creation logic here
  addNode(nodeData)
}

</script>

<template>
  <Background />
  <p style="color: white; font-size: 60px; text-align: center; margin-top: 200px;" v-if="!ready">Loading...</p>
  <div class="base-flow-container" v-if="ready">
    <VueFlow :nodes="nodes" :edges="edges">
      
    </VueFlow>
    <MenuBar 
      @add-node="handleAddNode"
      @reset="resetFlowState"
    />
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