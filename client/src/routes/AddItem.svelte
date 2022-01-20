<script>
    import ItemForm from "../components/ItemForm.svelte";
    import Group from "./Group.svelte";

    let title = "";
    let desc = "";
    let count = 1;
    let tags = [];
    let group;

    $: console.log(group);
    $: tags = tagsStr.split(" ").filter((x) => x.length);
    let tagsStr = tags.length ? tags.join(" ") : "";

    const formSubmit = async (event) => {
        event.preventDefault();
        if (!title) return alert("title cannot be blank");
        const res = await fetch("/api/add-item", {
            method: "POST",
            body: JSON.stringify({ title, desc, count, tags }),
        });
        console.log(res);
        res.json().then((ress) => {
            console.log(group);
            let id;
            if (ress.err) alert(ress.err);
            else ({ id } = ress);
            if (!group.id || group.id === -1) {
                if (!res.err) {
                    location.hash = `#/item/${encodeURIComponent(title)}`;
                }
                return;
            } else if (id) {
                fetch(`/api/add-item-to-group?group=${group.id}&item=${id}`, {
                    method: "POST",
                })
                    .then((res) => {
                        console.log(res);
                        if (res.status === 200) {
                            alert(`item added to ${group.title} successfully`);
                            location.hash = `#/item/${encodeURIComponent(
                                title
                            )}`;
                        } else return res.json();
                    })
                    .then((res) => alert(res.err));
            }
        });
        if (res.status === 200) {
            alert("item added successfully");
        }
    };
</script>

<div>
    <h1>Add Item</h1>
    <ItemForm
        bind:count
        bind:desc
        bind:title
        bind:tagsStr
        bind:group
        btnText="Add Item"
        {formSubmit}
    />
</div>

<style>
    div {
        display: flex;
        flex-direction: column;
        width: 80%;
    }
</style>
