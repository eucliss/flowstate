<script setup lang="ts">
import { ref, watch, onMounted, provide } from 'vue'
import type { Node, Edge } from '@vue-flow/core'  
import { VueFlow, useVueFlow } from '@vue-flow/core'
import { Controls } from '@vue-flow/controls'
import Background from './components/Background.vue'
// import { Background } from '@vue-flow/background'
import QueryNode from './components/nodes/QueryNode.vue'
import TextNode from './components/nodes/TextNode.vue'
import FlowNode from './components/nodes/FlowNode.vue'
import MenuBar from './components/MenuBar.vue'
import NodeDrawer from './components/NodeDrawer.vue'
import { testing, nodes, edges, getFlowState, updateNode, connectEdge, addNode, resetFlowState, updateAllNodesStatus, handleEdgeKeyDown, handleNodeKeyDown} from './functions'
import type { ComparisonType, globalCustomNodeTypes, globalCustomNode } from './functions'
import CountNode from './components/nodes/CountNode.vue'
import useDragAndDrop from './useDnD'
import Sidebar from './components/Sidebar.vue'


const { onDragOver, onDrop, onDragLeave, isDragOver } = useDragAndDrop()

type CustomNodeTypes = globalCustomNodeTypes
type CustomNode = globalCustomNode

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
  getSelectedEdges,
  onNodeDragStop,
  onConnect,
  onNodeClick
} = instance

onNodeDragStop(async (event) => {
  // If the node is dropped on the pane, ignore it
  console.log("node dragged: ", event)
  const updated_node = {
    id: event.node.id,
    position: {
      x: Math.floor(event.node.position.x),
      y: Math.floor(event.node.position.y),
    },
    data: {
      label: event.node.data.label,
      sql: event.node.data.sql,
      successRoute: event.node.data.successRoute,
      failureRoute: event.node.data.failureRoute,
      status: event.node.data.status,
      count: event.node.data.count,
      content: event.node.data.content,
    },
    type: event.node.type,
  }
  nodes.value.push(updated_node)
  await updateNode(updated_node)
})

onConnect(async (event) => {
  console.log("edge connected: ", event)
  const id = crypto.randomUUID()
  const new_edge = {
    id: id,
    source: event.source,
    target: event.target,
    animated: false,
    sourceHandle: event.sourceHandle,
    targetHandle: event.targetHandle,
    type: "smoothstep",
    // updatable: true,
    // label: "test",
  }
  edges.value.push(new_edge)
  await connectEdge(new_edge)
})

const localAddNode  = async (nodeData: { name: string, sql: string, type: string, successRoute: ComparisonType, failureRoute: ComparisonType }) => {
  console.log("Adding node: ", nodeData)
  const id = crypto.randomUUID()

  const new_node: Node = {
    id: id,
    type: nodeData.type,
    position: {
      x: 0,
      y: 100,
    },
    data: {
      label: nodeData.name,
      sql: nodeData.sql,
      successRoute: nodeData.successRoute,
      failureRoute: nodeData.failureRoute,
      count: 0,
      status: false,
      content: "",
    },
  }
  nodes.value.push(new_node)
  console.log("Nodes: ", nodes.value)
  await addNode(new_node)
  await getFlowState()
}

function removeEdge(id) {
  edges.value = edges.value.filter((edge) => edge.id !== id)
}

onMounted(async () => {
  await getFlowState()
  console.log("Mounted nodes", nodes.value)
  console.log("Mounted edges", edges.value)
  await updateAllNodesStatus()
  ready.value = true
  window.addEventListener('keydown', handleKeyDown)
  // const i = setInterval(async () => {
  //   await updateAllNodesStatus()
  // }, 5000)
})

const handleKeyDown = async (event) => {
  // Ignore delete/backspace if we're in a text input/textarea
  if ((event.key === 'Delete' || event.key === 'Backspace') && 
      (event.target.tagName === 'INPUT' || event.target.tagName === 'TEXTAREA')) {
    return;
  }

  console.log("handleKeyDown pressed")
  console.log("key pressed: ", event.keyCode)
  console.log("selected edges: ", getSelectedEdges.value)
  console.log("selected nodes: ", getSelectedNodes.value)
  if (getSelectedEdges.value.length > 0) {
    await handleEdgeKeyDown(event, getSelectedEdges.value)
  }
  if(getSelectedNodes.value.length > 0) {
    const connections = getConnectedEdges(getSelectedNodes.value[0].id)
    await handleNodeKeyDown(event, getSelectedNodes.value, connections)
    if (event.key == 'Delete' || event.key == 'Backspace') {
      onPaneClick()
    }
  }
}

const resetFlow = async () => {
  await resetFlowState()
  await getFlowState()
}


const selectedNode = ref<Node | null>(null)
const drawerOpen = ref(false)
onNodeClick((event) => {
  onPaneClick()
  selectedNode.value = event.node
  drawerOpen.value = true  // Open drawer when clicking on a node
  console.log("selected node: ", selectedNode.value) 
})

// Add click handler for the vue-flow pane
const onPaneClick = () => {
  drawerOpen.value = false  // Close drawer when clicking on the pane
  selectedNode.value = null
}

const updateNodeData = async (nodeId: string, newData: any) => {
  const updatedNode = {
    ...nodes.value.find(node => node.id === nodeId),
    data: newData
  }
  await updateNode(updatedNode)
}

provide('selectedNode', selectedNode)
provide('updateNodeData', updateNodeData)

onPaneReady((i) => i.fitView())
</script>

<template>
  <Controls class="controls"/>

  <p style="color: white; font-size: 60px; text-align: center; margin-top: 200px;" v-if="!ready">Loading...</p>
  <div class="base-flow-container dnd-flow" v-if="ready" @drop="onDrop">
    <Sidebar />
    
    <VueFlow 
      :nodes="nodes" 
      :edges="edges"
      @paneClick="onPaneClick"
      @dragover="onDragOver"
      @drop="onDrop"
      @dragleave="onDragLeave"
    >
      <Background/>

      <template #node-flowNode="flowNodeProps">
        <FlowNode v-bind="flowNodeProps"  />
      </template>
      
      <template #node-queryNode="queryNodeProps">
        <QueryNode v-bind="queryNodeProps" />
      </template>      
      <template #node-countNode="countNodeProps">
        <CountNode v-bind="countNodeProps" />
      </template>
      <template #node-textNode="textNodeProps">
        <TextNode 
          v-bind="textNodeProps" 
          :updateNodeData="updateNodeData"
        />
      </template>
    </VueFlow>
    <!-- <MenuBar 
      @add-node="localAddNode"
      @reset="resetFlowState"
    /> -->
    <NodeDrawer 
      v-if="drawerOpen"
      @closeDrawer="drawerOpen = false"
      :selectedNode="selectedNode"
      type="update"
    />
  </div>
</template>

<style>
/* import the necessary styles for Vue Flow to work */
@import '@vue-flow/core/dist/style.css';
@import '@vue-flow/core/dist/theme-default.css';
@import '@vue-flow/controls/dist/style.css';

/* Reposition controls to bottom right */
.controls {
  position: fixed !important;
  bottom: 20px !important;
  right: 20px !important;
  top: unset !important;
  left: unset !important;
  z-index: 5 !important;
}

.vue-flow__controls-button {
  background: #222222;
  border: 1px solid rgba(255, 255, 255, 0.12);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
  color: rgba(255, 255, 255, 0.7);
  transition: all 0.2s ease;
  width: 32px;  /* Slightly larger buttons */
  height: 32px;
}

.vue-flow__controls-button:hover {
  background: #2d2d2d;
  border-color: rgba(255, 255, 255, 0.3);
  box-shadow: 
    0 4px 12px rgba(0, 0, 0, 0.4),
    0 0 0 1px rgba(255, 255, 255, 0.1);
  color: rgba(255, 255, 255, 1);
  transform: translateY(-2px);
}

.vue-flow__controls-button:active {
  transform: translateY(0px);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.3);
}

/* Add a subtle glow to the controls container on hover */
.controls:hover .vue-flow__controls-button {
  border-color: rgba(255, 255, 255, 0.2);
}

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
        </template> -->