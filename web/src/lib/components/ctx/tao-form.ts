import {createContext} from "$lib/context";
import type {Writable} from "svelte/store";

export const ctx = createContext<Writable<boolean>>();
