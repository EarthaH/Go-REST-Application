import { Request, Response, NextFunction } from 'express';
import axios, { AxiosResponse } from 'axios';

interface File {
    fileName: String;
}

const getHomePage = async (req: Request, res: Response, next: NextFunction) => {
    let result: AxiosResponse = await axios.get(`http://localhost:8080/`);
    let greeting: String = result.data;
    return res.status(200).json({
        message: greeting
    });
};

const getFiles = async (req: Request, res: Response, next: NextFunction) => {
    let result: AxiosResponse = await axios.get(`http://localhost:8080/files/`);
    let files: [File] = result.data;
    return res.status(200).json({
        message: files
    });
};

const makeFile = async (req: Request, res: Response, next: NextFunction) => {
    let fname: string = req.params.filename;
    
    let response: AxiosResponse = await axios.get(`http://localhost:8080/files/new/${fname}`);
    return res.status(200).json({
        message: response.data
    });
};

const getFile = async (req: Request, res: Response, next: NextFunction) => {
    let fname: string = req.params.filename;
    
    let result: AxiosResponse = await axios.get(`http://localhost:8080/files/${fname}`);
    let file: File = result.data;
    return res.status(200).json({
        message: file
    });
};

const updateFile = async (req: Request, res: Response, next: NextFunction) => {
    let fname: string = req.params.filename;
    let lines: string = req.body.lines;
    
    let response: AxiosResponse = await axios.post(`http://localhost:8080/files/${fname}/save`, {
        "line": lines
    })
    return res.status(200).json({
        message: response.data
    });
};

const renameFile = async (req: Request, res: Response, next: NextFunction) => {
    let oldname: string = req.params.oldname;
    let newname: string = req.params.newname;
    
    let response: AxiosResponse = await axios.get(`http://localhost:8080/files/replace/${oldname}/${newname}`);
    return res.status(200).json({
        message: response.data
    });
};

const deleteFile = async (req: Request, res: Response, next: NextFunction) => {
    let fname: string = req.params.filename;
    
    let response: AxiosResponse = await axios.get(`http://localhost:8080/files/delete/${fname}`);
    return res.status(200).json({
        message: response.data
    });
};

export default { getHomePage, getFiles, getFile, updateFile, renameFile, deleteFile };
