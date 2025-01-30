import type { Node, Edge } from '@vue-flow/core'  
import { ref } from 'vue'
export const URL = 'http://localhost:3000'

// SELECT * FROM 'test_torq3'

export const nodeTypes = ['default', 'custom']

export type ComparisonType = {
    leftValue: string
    rightValue: string
    operator: string
}

export type NodeData = {
    label: string
    sql: string
    successRoute: ComparisonType
    failureRoute: ComparisonType
}

export type EdgeData = {
    id: string
    source: string
    target: string
    animated: boolean
}

export const nodes = ref<Node[]>([])
export const edges = ref<Edge[]>([])

export const convertNodetoGoNode = (node: Node) => {
    console.log("Converting node to Go node: ", node)
    return {
        id: node.id,
        type: node.type,
        x: node.position.x,
        y: node.position.y,
        label: node.data.label,
        sql: node.data.sql,
    }
}

export const convertGoNodeToNode = (n: { Id: string, Type: string, X: number, Y: number, Label: string, SQL: string }) => {
    console.log("Converting Go node to node: ", n)
    return {
        id: n["Id"],
        type: n["Type"],
        position: {
            x: n["X"],
            y: n["Y"],
        },
        data: {
            label: n["Label"],
            sql: n["SQL"],
        },
    }
}

export const convertGoEdgeToEdge = (e: { Id: string, Source: string, Target: string, Animated: boolean }) => {
    console.log("Converting Go edge to edge: ", e)
    return {
        id: e["Id"],
        source: e["Source"],
        target: e["Target"],
        animated: e["Animated"],
    }
}

export const convertEdgeToGoEdge = (e: Edge) => {
    console.log("Converting edge to Go edge: ", e)
    return {
        Id: e.id,
        Source: e.source,
        Target: e.target,
        Animated: e.animated,
    }
}

export const getFlowState = async () => {
    console.log("Loading flow state ... ")
    const response = await fetch(`${URL}/load-flow-state`, {
        method: 'GET',
    })
    const data = await response.json()

    nodes.value = data.nodes.map(convertGoNodeToNode)
    edges.value = data.edges.map(convertGoEdgeToEdge)
    console.log("Flow state loaded: ", nodes)
}

export const updateNode = async (node: Node) => {
    const response = await fetch(`${URL}/update-node`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(convertNodetoGoNode(node)),
    })
    const data = await response.json()
    return data
}

export const addNode = async (node: Node) => {
    const goNode = convertNodetoGoNode(node)
    console.log("Adding node to Go: ", goNode)
    const response = await fetch(`${URL}/add-node`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(goNode),
    })
    const data = await response.json()
    return data
}

export const connectEdge = async (edge: Edge) => {
    const response = await fetch(`${URL}/connect-edge`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(edge),
    })
    const data = await response.json()
    return data
}

export const resetFlowState = async () => {
    const response = await fetch(`${URL}/reset`, {
        method: 'GET',
    })
    const data = await response.json()
    return data
}

export const testing = () => {
    console.log('testing')
}