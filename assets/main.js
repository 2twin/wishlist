window.onload = () => {
    const openModalButtons = document.querySelectorAll(".open-modal");
    const overlay = document.querySelector(".overlay");

    openModalButtons.forEach(btn => {
        const id = btn.dataset.modal;
        const modal = document.getElementById(id);

        btn.addEventListener("click", async () => {
            if (!modal) {
                return
            }
            overlay.classList.add("active");
            modal.classList.add("active");

            switch (btn.dataset.modal) {
                case "edit-wish": {
                    const req = {
                        username: btn.closest(".user").dataset.username,
                        wish_id: Number.parseInt(btn.closest(".item").dataset.id)
                    }

                    const response = await fetch("https://wishlist-v9g5wojd.b4a.run/get_wish", {
                        method: "POST",
                        body: JSON.stringify(req),
                        headers: {
                            'Accept': 'application/json',
                            'Content-Type': 'application/json'
                        },
                    })

                    const data = await response.json();

                    modal.querySelector("#input-wish-title").setAttribute("value", data.Title);
                    modal.querySelector("#input-wish-link").setAttribute("value", data.Link);

                    const form = modal.querySelector(".modal__form")
                    form.addEventListener("submit", async (e) => {
                        const req = {
                            username: btn.closest(".user").dataset.username,
                            wish_id: data.ID,
                            title: modal.querySelector("#input-wish-title").value,
                            link: modal.querySelector("#input-wish-link").value
                        }

                        await fetch("https://wishlist-v9g5wojd.b4a.run/edit_wish", {
                            method: "PATCH",
                            body: JSON.stringify(req),
                            headers: {
                                'Accept': 'application/json',
                                'Content-Type': 'application/json'
                            },
                        })
                    })
                }

                case "add-wish": {
                    const form = modal.querySelector(".modal__form")
                    form.addEventListener("submit", async (e) => {
                        const req = {
                            username: btn.closest(".user").dataset.username,
                            title: modal.querySelector("#input-wish-title").value,
                            link: modal.querySelector("#input-wish-link").value
                        }

                        await fetch("https://wishlist-v9g5wojd.b4a.run/add_wish", {
                            method: "POST",
                            body: JSON.stringify(req),
                            headers: {
                                'Accept': 'application/json',
                                'Content-Type': 'application/json'
                            },
                        })
                    })
                }

                case "remove-wish": {
                    const form = modal.querySelector(".modal__form")
                    form.addEventListener("submit", async (e) => {
                        const req = {
                            username: btn.closest(".user").dataset.username,
                            wish_id: Number.parseInt(btn.closest(".item").dataset.id),
                        }

                        console.log(req)

                        await fetch("https://wishlist-v9g5wojd.b4a.run/remove_wish", {
                            method: "DELETE",
                            body: JSON.stringify(req),
                            headers: {
                                'Accept': 'application/json',
                                'Content-Type': 'application/json'
                            },
                        })
                    })
                }
            }
        })
    })

    const modalSubmitButtons = document.querySelectorAll(".modal__btn");
    modalSubmitButtons.forEach(btn => {
        btn.addEventListener("click", () => {

        })
    })

    const completeWishButtons = document.querySelectorAll(".btn-complete")
    completeWishButtons.forEach(btn => {
        btn.addEventListener("click", async () => {
            const req = {
                username: btn.closest(".user").dataset.username,
                wish_id: Number.parseInt(btn.closest(".item").dataset.id)
            }

            await fetch("https://wishlist-v9g5wojd.b4a.run/toggle_wish_status", {
                method: "PATCH",
                body: JSON.stringify(req),
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                },
            })

            btn.closest(".item").classList.toggle("completed");
        })
    })

    const closeModalButtons = document.querySelectorAll(".modal__close");
    closeModalButtons.forEach(btn => {
        btn.addEventListener("click", () => {
            document.querySelector(".overlay").classList.remove("active");
            document.querySelector(".modal.active").classList.remove("active");
        })
    })
}