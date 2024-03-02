import "./style.css";

import {
  GetMods,
  OpenFileDialog,
  AddMod,
  GetGameDir,
  OpenFolderDialog,
  SetGameDir,
  RemoveMod,
} from "../wailsjs/go/main/App";
import { EventsOn } from "../wailsjs/runtime/runtime";
import { addModElement, showErrorModal, showFormModal, showSettingsModal } from "./elements";

const logs = document.querySelector<HTMLTextAreaElement>("#logOutput");

let MODS: string[] = [];

window.openFileDialog = async function () {
  const file = await OpenFileDialog();

  if (!file.length) {
    return;
  }

  const fileName = file.split("\\").pop() || file.split("/").pop() || "";
  showFormModal(fileName.split(".").shift() || fileName, file);
};

window.openFolderDialog = async function () {
  return await OpenFolderDialog();
};

window.openSettingsModal = async function () {
  const gameDir = await GetGameDir();
  showSettingsModal(gameDir);
};

window.addMod = async function (name: string, path: string) {
  await AddMod(name, path);
};

window.removeMod = async function (name: string) {
  await RemoveMod(name);
};

window.saveFolder = async function (folder: string) {
  await SetGameDir(folder);
};

EventsOn("log:new", (data: string) => {
  const log = document.createElement("div");
  log.innerHTML = data;
  logs!.appendChild(log);
  scrollToBottom();
});

EventsOn("error:new", (data: string) => {
  const log = document.createElement("div");
  log.innerHTML = `<span class="text-red-500">${data}</span>\n`;
  logs!.appendChild(log);
  scrollToBottom();

  showErrorModal(data);
});

EventsOn("success:new", (data: string) => {
  const log = document.createElement("div");
  log.innerHTML = `<span class="text-green-500">${data}</span>\n`;
  logs!.appendChild(log);
  scrollToBottom();
});

EventsOn("mods:refresh", () => {
  refreshMods();
});

document.querySelector<HTMLInputElement>("#searchMods")?.addEventListener("input", (e) => {
  e.preventDefault();

  const input = e.target as HTMLInputElement;
  const value = input.value;

  const filtered = MODS.filter((mod) => mod.toLowerCase().includes(value.toLowerCase()));
  showMods(filtered);
});

function showMods(mods: string[]) {
  const modList = document.querySelector<HTMLUListElement>("#modList");
  if (!modList) return;
  modList!.innerHTML = "";

  for (const mod of mods) {
    addModElement(modList, mod);
  }
}

async function refreshMods() {
  const mods = await GetMods();
  MODS = mods.map((mod) => mod.Name);
  showMods(MODS);
}

function scrollToBottom() {
  const logOutput = document.querySelector("#logOutput");
  logOutput?.scrollTo(0, logOutput.scrollHeight);
}

refreshMods();

declare global {
  interface Window {
    openFileDialog: () => void;
    openFolderDialog: () => Promise<string>;
    addMod: (name: string, path: string) => void;
    removeMod: (name: string) => void;
    saveFolder: (folder: string) => void;
    openSettingsModal: () => void;
  }
}
