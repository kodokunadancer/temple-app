import { ActionTree, ActionContext } from 'vuex'

export const name = 'auth'
export const namespaced = true

export interface User {
    name: string
    email: string
    password: string
}

export interface State {
    user: User | null
}

export interface Actions<S, R> extends ActionTree<S, R> {
    setUser(context: ActionContext<S, R>, param: User): void
}

export const types = {
    SET_USER: 'SET_USER'
}