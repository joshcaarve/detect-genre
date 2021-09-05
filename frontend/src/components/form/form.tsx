import * as React from 'react';
import { useState, useEffect } from 'react';
import axios from 'axios';
import './form.css';
interface Props {
    name: string
    value: string
}

const FormComponent: React.FC<Props> = props => {

    const handleSubmit = async (event: React.SyntheticEvent) => {
        const target = event.target as typeof event.target & {
            operation: { value: string };
            genre: { value: string };
            parentgenre: { value: string };
        };
        //console.log(target)

        if (target.operation.value === "add") {
            await axios({
                method: "post",
                url: "api/list",
                data: {
                    genre: target.genre.value,
                    parentGenre: target.parentgenre.value,
                }
            });
        } else if (target.operation.value === "delete") {
            const url = "api/list/" + target.genre.value;
            await axios({
                method: "delete",
                url: url
            });
        }
    }

    return (
        <div>
            <form onSubmit={handleSubmit}>
                <label>
                    Operation
                    <select id="operation" name="operation">
                        <option value="add">Add</option>
                        <option value="delete">Delete</option>
                    </select>
                </label>
                <label>
                    Genre:
                    <input type="text" name="genre" />
                </label>
                <label>
                    Parent Genre:
                    <input type="text" name="parentgenre" />
                </label>
                <input type="submit" value="Submit" />
            </form >
        </div >
    )
}

export default FormComponent;
