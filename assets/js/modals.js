// Performs basic AJAX functionality to fetch and render a modal.

function showLoadingOverlay() {
    $(`#loadingOverlay`)[0].classList.remove("hidden");
}

function hideLoadingOverlay() {
    $(`#loadingOverlay`)[0].classList.add("hidden");
}

function closeModal() {
    $(`#dynamicModal`)[0].classList.add('modal-closing');
    setTimeout(() => {
        $("#modalContainer").empty();
    }, 200);
}

function openModal(id) {
    showLoadingOverlay(); // Show loading overlay before fetching content
    $("#modalContainer").load(`/modal?id=${id}`, function(responseTxt, statusTxt, xhr){
        hideLoadingOverlay(); // Hide loading overlay after content is fetched
        if (statusTxt == "error") {
            console.log("Error loading modal: " + xhr.status + ": " + xhr.statusText);
        }
    });
}

// Open a modal when a button is clicked
document.querySelectorAll(".open-modal-button").forEach(button => {
    button.addEventListener("click", () => {
        openModal(button.dataset.itemId); // Use the item ID from the button instead of the target, as it tends to be undefined
    });
});