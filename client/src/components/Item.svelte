<script>
    import ItemForm from "./ItemForm.svelte";
    import { duplicate } from "../utils";

    export let title;
    export let desc;
    export let count;
    export let tags;
    export let id;
    export let created_at;
    export let updated_at;

    let createdAt = new Date(created_at).toLocaleString();

    let tagsStr = tags.length > 0 ? tags.join(" ") : "";

    $: titleValue = title;
    $: descValue = desc;
    $: countValue = count;
    $: tagsValue = tagsStr.split(" ").filter((x) => x.length > 0);

    let group = { id: -1, title: "" };
    let groupValue = group;
    $: console.log(group);

    fetch(`/api/group-of-item?id=${id}`)
        .then((res) => res.json())
        .then((res) => {
            console.log(res);
            if (res) {
                if (res.err);
                else {
                    group = { id: res.id, title: res.title };
                    groupValue = group;
                }
            }
        });

    const formSubmit = (event) => {
        event.preventDefault();
        if (!title) return alert("title cannot be blank");
        const dup = duplicate(tagsValue);
        if (dup) return alert(`cannot have more than one "${dup}" tag`);
        fetch(`/api/update-item?id=${id}`, {
            method: "POST",
            body: JSON.stringify({
                title: titleValue,
                desc: descValue,
                count: countValue,
                tags: tagsValue,
            }),
        })
            .then((res) => {
                console.log(res);
                editMode = false;
                if (res.status === 200) return res.json();
            })
            .then((res) => {
                console.log(res);
                if (res) {
                    if (typeof res.tags === "undefined") res.tags = [];
                    ({ title, desc, count, tags, created_at, updated_at } =
                        res);
                }
                titleValue = title;
                descValue = desc;
                countValue = count;
                tagsValue = tags;
                tagsStr = tags.length > 0 ? tags.join(" ") : "";
                window.location.hash = "#/item/" + encodeURIComponent(title);
            })
            .catch(console.log);
        if (groupValue.id === -1 && group.id !== -1) {
            fetch(`/api/remove-item-from-group?item=${id}`, {
                method: "POST",
            })
                .then((res) => {
                    console.log(res);
                    if (res.status === 200) {
                        alert(`item removed  from ${group.title} successfully`);
                        group = { id: -1, title: "" };
                    } else return res.json();
                })
                .then((res) => {
                    if (res.err) alert(res.err);
                })
                .catch(console.log);
        } else if (group.title !== groupValue.title) {
            fetch(`/api/add-item-to-group?group=${groupValue.id}&item=${id}`, {
                method: "POST",
            })
                .then((res) => {
                    console.log(res);
                    if (res.status === 200) {
                        alert(`item added to ${groupValue.title} successfully`);
                        group = groupValue;
                    } else return res.json();
                })
                .then((res) => {
                    if (res.err) alert(res.err);
                })
                .catch(console.log);
        }
    };

    const deleteItem = () => {
        if (confirm(`Are you sure you want to delete ${title}`))
            fetch(`/api/delete-item?id=${id}`, {
                method: "DELETE",
            });
        location.hash = "#";
    };

    let editMode = false;
</script>

{#if editMode}
    <ItemForm
        bind:title={titleValue}
        bind:desc={descValue}
        bind:count={countValue}
        bind:group={groupValue}
        bind:tagsStr
        btnText="update"
        {formSubmit}
    />
{:else}
    <h2>{title}</h2>
    <p>created at: {createdAt}</p>
    <p>description: {desc}</p>
    <p>group: {group.title && group.title}</p>
    <p>count: {count}</p>
    <ul>
        <p>tags:</p>
        {#each tags as tag}
            <li><a href={`#/tag/${tag}`}>{tag}</a></li>
        {/each}
    </ul>
    <button
        on:click={() => {
            editMode = true;
        }}>edit</button
    >
    <button class="deleteBtn" on:click={deleteItem}>delete</button>
{/if}

<style>
    ul {
        padding: 0;
    }
    li {
        list-style: none;
        display: inline;
        margin: 0 4px;
    }

    button {
        max-width: 320px;
    }

    .deleteBtn {
        color: #eee;
        background-color: crimson;
    }
</style>
