<script setup lang="ts">
import { ref, computed } from 'vue'
import { nodeTypes } from '../functions'
import type { Node } from '@vue-flow/core'
import type { PropType } from 'vue'
import Comparison from './Comparison.vue'
import type { ComparisonType} from '../functions'
import { addNode, updateNode } from '../functions'

const emit = defineEmits(['closeDrawer'])
const nodeName = ref('')
const sqlQuery = ref('')
const type = ref(nodeTypes[0])
const jsonData = ref('')
const nodeType = ref('')

const successRoute = ref<ComparisonType>({
    leftValue: "",
    rightValue: "",
    operator: ""
})
const failureRoute = ref<ComparisonType>({
    leftValue: "",
    rightValue: "",
    operator: ""
})

const props = defineProps({
  selectedNode: {
    type: Object as PropType<Node>,
    required: false
  },
  type: {
    type: String,
    required: false
  }
})

const addNewNode = () => {
    console.log('AddNodeDrawer - Emitting with name:', nodeName.value, 'SQL:', sqlQuery.value)
    console.log('AddNodeDrawer - Success route:', successRoute.value)
    console.log('AddNodeDrawer - Failure route:', failureRoute.value)
    const new_node: Node = {
        id: crypto.randomUUID(),
        type: type.value,
        position: {
            x: 0,
            y: 100,
        },
        data: {
            label: nodeName.value,
            sql: sqlQuery.value,
            successRoute: successRoute.value,
            failureRoute: failureRoute.value,
        },
    }
    addNode(new_node)
    nodeName.value = ''
    sqlQuery.value = ''
    jsonData.value = ''
    type.value = nodeTypes[0]
    emit('closeDrawer')
}
const updateExistingNode = () => {
    console.log('AddNodeDrawer - Emitting with name:', nodeName.value, 'SQL:', sqlQuery.value)
    const updated_node: Node = {
        id: props.selectedNode.id,
        type: type.value,
        position: {
            x: props.selectedNode.position.x,
            y: props.selectedNode.position.y,
        },
        data: {
            label: nodeName.value,
            sql: sqlQuery.value,
            successRoute: successRoute.value,
            failureRoute: failureRoute.value,
        },
    }
    updateNode(updated_node)
    nodeName.value = ''
    sqlQuery.value = ''
    jsonData.value = ''
    type.value = nodeTypes[0]
    emit('closeDrawer')
}

const handleSubmit = () => {
  if (props.type === 'add') {
    addNewNode()
  } else {
    updateExistingNode()
  }

}

const handleTest = async () => {
  console.log('AddNodeDrawer - Testing JSON data:', sqlQuery.value)

  const response = await fetch('http://localhost:3000/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },

    body: JSON.stringify({
      query: sqlQuery.value,
      limit: 10,
      start: 1738040675782000,
      end: 1738041575782000,
      sourceType: 'flowstate'
    }),
  })
  const data = await response.json()
  console.log('AddNodeDrawer - Test response:', data)
  jsonData.value = JSON.stringify(data[0], null, 2)
}

if (props.selectedNode) {
  console.log("selected node: ", props.selectedNode)
  nodeName.value = props.selectedNode.data.label
  sqlQuery.value = props.selectedNode.data.sql
  type.value = props.selectedNode.type
  jsonData.value = props.selectedNode.data.json
  try {
    successRoute.value = props.selectedNode.data.successRoute
    failureRoute.value = props.selectedNode.data.failureRoute
  } catch (error) {
    console.error('Error parsing successRoute or failureRoute:', error)
  }

  
  console.log('Selected node type:', type.value)
  console.log('Full node data:', props.selectedNode)
}
console.log('Available nodeTypes:', nodeTypes)
console.log('Current selected type:', type.value)

const updateSuccessRoute = (value: ComparisonType) => {
  successRoute.value = value
  console.log('AddNodeDrawer - Updated successRoute:', successRoute.value)
}

const updateFailureRoute = (value: ComparisonType) => {
  failureRoute.value = value
  console.log('AddNodeDrawer - Updated failureRoute:', failureRoute.value)
}

</script>

<template>
    <Transition name="drawer">
      <div class="side-drawer">
        <div class="drawer-content">
          <div class="drawer-header">
            <h3>{{ props.type === 'add' ? 'Add New Node' : 'Update Node' }}</h3>
            <button class="close-button" @click="$emit('closeDrawer')">×</button>
          </div>
          
          <div class="drawer-body">
            <div class="input-row">
              <div class="input-group">
                <label for="nodeName">Node Name</label>
                <input 
                  id="nodeName"
                  v-model="nodeName" 
                  placeholder="Enter node name"
                  type="text"
                />
              </div>

              <div class="input-group">
                <label for="nodeType">Type</label>
                <select 
                  id="nodeType"
                  v-model="type"
                  class="type-select"
                >
                  <option v-for="t in nodeTypes" :key="t" :value="t">
                    {{ t }}
                  </option>
                </select>
              </div>
            </div>
  
            <div class="input-group">
              <label for="sqlQuery">SQL Query</label>
              <textarea
                id="sqlQuery"
                v-model="sqlQuery"
                placeholder="Enter SQL query"
                class="sql-input"
                rows="1"
              ></textarea>
            </div>

            <div class="input-group">
              <div class="label-row">
                <label for="jsonData">JSON Data</label>
                <button class="test-button" @click="handleTest">
                  <i class="fas fa-vial"></i>
                  Test
                </button>
              </div>
              <textarea
                id="jsonData"
                v-model="jsonData"
                placeholder="Enter JSON data"
                class="json-input"
                rows="6"
              ></textarea>
            </div>

            <div class="input-group">
              <label>Success Route</label>
              <Comparison
                :model-value="successRoute"
                @update:model-value="updateSuccessRoute"
              />
              <label>Failure Route</label>
              <Comparison 
                :model-value="failureRoute"
                @update:model-value="updateFailureRoute"
              />
            </div>
          </div>
  
          <div class="drawer-footer">
            <button class="btn-primary" @click="handleSubmit">{{ props.type === 'add' ? 'Add Node' : 'Update Node' }}</button>
            <button class="btn-secondary" @click="$emit('closeDrawer')">Cancel</button>
          </div>
        </div>
      </div>
    </Transition>
  </template>

<style scoped>
/* Base drawer styles */
.side-drawer {
  position: fixed;
  top: 0;
  right: 0;
  width: 40%;
  height: 100vh;
  background: #1e1e1e;
  border-left: 1px solid #333;
  box-shadow: -8px 0 32px rgba(0, 0, 0, 0.2);
  z-index: 1000;
}

.drawer-content {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.drawer-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 24px;
  border-bottom: 1px solid #333;
}

.drawer-header h3 {
  color: #fff;
  margin: 0;
  font-size: 1.5rem;
  font-weight: 500;
}

/* Body and columns styling */
.drawer-body {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow-y: auto;
  padding: 24px;
}

.input-row {
  display: flex;
  gap: 20px;
  margin-bottom: 24px;
}

.input-row .input-group {
  margin-bottom: 0;
  flex: 1;
}

.input-group {
  margin-bottom: 24px;
}

.columns-container {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 20px;
  flex: 1;
  min-height: 0;
}

.column {
  background: #2d2d2d;
  border-radius: 8px;
  padding: 16px;
  border: 1px solid #404040;
  display: flex;
  flex-direction: column;
  height: 100%;
}

.column h4 {
  color: #fff;
  margin: 0 0 16px 0;
  font-size: 1rem;
  font-weight: 500;
}

.column-content {
  flex: 1;
}

/* Input styling */
.input-group label {
  display: block;
  color: #999;
  margin-bottom: 8px;
  font-size: 0.9rem;
}

input {
  width: 100%;
  padding: 12px 16px;
  background: #2d2d2d;
  border: 1px solid #404040;
  border-radius: 6px;
  color: #fff;
  font-size: 1rem;
  transition: all 0.2s ease;
}

input:focus {
  outline: none;
  border-color: #0066cc;
  box-shadow: 0 0 0 2px rgba(0, 102, 204, 0.2);
}

input::placeholder {
  color: #666;
}

.sql-input {
  width: 100%;
  padding: 12px 16px;
  background: #2d2d2d;
  border: 1px solid #404040;
  border-radius: 6px;
  color: #fff;
  font-size: 1rem;
  transition: all 0.2s ease;
  resize: vertical;
  min-height: 150px;
  font-family: monospace;
}

.sql-input:focus {
  outline: none;
  border-color: #0066cc;
  box-shadow: 0 0 0 2px rgba(0, 102, 204, 0.2);
}

.sql-input::placeholder {
  color: #666;
}

.json-input {
  width: 100%;
  padding: 12px 16px;
  background: #2d2d2d;
  border: 1px solid #404040;
  border-radius: 6px;
  color: #fff;
  font-size: 1rem;
  transition: all 0.2s ease;
  resize: vertical;
  min-height: 200px;
  font-family: monospace;
}

.json-input:focus {
  outline: none;
  border-color: #0066cc;
  box-shadow: 0 0 0 2px rgba(0, 102, 204, 0.2);
}

.json-input::placeholder {
  color: #666;
}

/* Footer styling */
.drawer-footer {
  padding: 24px;
  border-top: 1px solid #333;
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

/* Button styles */
.close-button {
  background: transparent;
  border: none;
  color: #666;
  font-size: 24px;
  cursor: pointer;
  padding: 4px 8px;
  border-radius: 4px;
  transition: all 0.2s ease;
}

.close-button:hover {
  color: #fff;
  background: rgba(255, 255, 255, 0.1);
}

.btn-primary,
.btn-secondary {
  padding: 10px 20px;
  border-radius: 6px;
  font-size: 0.95rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
}

.btn-primary {
  background: #0066cc;
  color: white;
  border: none;
}

.btn-primary:hover {
  background: #0077ee;
}

.btn-secondary {
  background: transparent;
  color: #999;
  border: 1px solid #404040;
}

.btn-secondary:hover {
  background: rgba(255, 255, 255, 0.1);
  color: #fff;
}

/* Animation */
.drawer-enter-active,
.drawer-leave-active {
  transition: transform 0.3s ease;
}

.drawer-enter-from,
.drawer-leave-to {
  transform: translateX(100%);
}

/* Add these new styles */
.type-select {
  width: 100%;
  padding: 12px 16px;
  background: #2d2d2d;
  border: 1px solid #404040;
  border-radius: 6px;
  color: #fff;
  font-size: 1rem;
  transition: all 0.2s ease;
  cursor: pointer;
}

.type-select:focus {
  outline: none;
  border-color: #0066cc;
  box-shadow: 0 0 0 2px rgba(0, 102, 204, 0.2);
}

.type-select option {
  background: #2d2d2d;
  color: #fff;
}

.label-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.test-button {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 6px 12px;
  background: #2d2d2d;
  border: 1px solid #404040;
  border-radius: 4px;
  color: #fff;
  font-size: 0.9rem;
  cursor: pointer;
  transition: all 0.2s ease;
}

.test-button:hover {
  background: #404040;
}

.test-button i {
  font-size: 0.9rem;
}

.comparison-container {
  display: flex;
  gap: 12px;
  align-items: center;
}

.comparison-select {
  padding: 12px 16px;
  background: #2d2d2d;
  border: 1px solid #404040;
  border-radius: 6px;
  color: #fff;
  font-size: 1rem;
  transition: all 0.2s ease;
  cursor: pointer;
}

.comparison-select:focus {
  outline: none;
  border-color: #0066cc;
  box-shadow: 0 0 0 2px rgba(0, 102, 204, 0.2);
}

.comparison-input {
  width: 120px;
}
</style> 