export function showSettingsModal(folder: string) {
  const body = document.querySelector("body");
  if (!body) return;

  const div = document.createElement("div");
  div.id = "popup-modal-settings";

  div.innerHTML = `
<div class="fixed inset-0 z-50 overflow-auto bg-smoke-light flex justify-center items-center">
  <div class="relative p-4 w-full max-w-md max-h-full">
    <div class="relative bg-white rounded-lg shadow dark:bg-gray-700">
      <div class="flex items-center justify-between p-4 md:p-5 border-b rounded-t dark:border-gray-600">
        <h3 class="text-lg font-semibold text-gray-900 dark:text-white">Add New Mod</h3>
        <button id="popup-modal-settings-close" type="button" class="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 ms-auto inline-flex justify-center items-center dark:hover:bg-gray-600 dark:hover:text-white" data-modal-toggle="crud-modal">
          <svg class="w-3 h-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 14">
            <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"/>
          </svg>
          <span class="sr-only">Close modal</span>
        </button>
      </div>
      <form class="p-4 md:p-5">
        <div class="grid gap-4 mb-4 grid-cols-2">
          <div class="col-span-2">
            <label for="folder" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Game folder</label>
            <input type="text" name="folder" id="folder" value="${folder}" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="Type mod name" required="">
          </div>
        </div>
        <button id="submit-select" type="button" class="text-white inline-flex items-center bg-green-700 hover:bg-green-800 focus:ring-4 focus:outline-none focus:ring-green-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800">
          Select folder
        </button>
        <button id="submit-save" type="button" class="text-white inline-flex items-center bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
          Save
        </button>
      </form>
    </div>
  </div>
</div> 
`;

  body.appendChild(div);
  const close = document.querySelector<HTMLButtonElement>("#popup-modal-settings-close");
  const select = document.querySelector<HTMLButtonElement>("#submit-select");
  const save = document.querySelector<HTMLButtonElement>("#submit-save");

  close?.addEventListener("click", () => {
    const modal = document.querySelector("#popup-modal-settings");
    modal?.remove();
  });

  select?.addEventListener("click", async () => {
    const folder = await window.openFolderDialog();
    if (folder) (document.querySelector("#folder") as HTMLInputElement).value = folder;
  });

  save?.addEventListener("click", async () => {
    const folder = (document.querySelector("#folder") as HTMLInputElement).value;
    await window.saveFolder(folder);

    const modal = document.querySelector("#popup-modal-settings");
    modal?.remove();
  });
}
