import { MutationTree, GetterTree } from 'vuex'
import { RootState } from '../../../store'
import { Actions, State, types, User } from './interface'
import axios from 'axios'

export const state = (): State => ({
    user: null
})

export const mutations: MutationTree<State> = {
    SET_USER(_state: State, user: User): void {
      _state.user = user
    }
}

export const getters: GetterTree<State, RootState> = {
    adminMode: (_state: State): State['user'] => _state.user
}

export const actions: Actions<State, RootState> = {
    setUser({ commit }, user: User): void {
      commit(types.SET_USER, user)
    }
}