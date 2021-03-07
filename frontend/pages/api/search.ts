import { SearchedObject } from "api";
import { NextApiRequest, NextApiResponse } from "next";
import { readFileSync } from 'fs'

export default (req: NextApiRequest, res: NextApiResponse) => {
    const data = readFileSync('/home/adam/Desktop/AKAI/frontend/pages/api/unnamed.jpg')

    const v: SearchedObject = {
        objectID: 1,
        title: 'Title',
        description: 'Description',
        img: data.toString('base64')
    };

    res.status(200).send(Array(10).fill(v));
}