/** source/routes/files.ts */
import express from 'express';
import controller from '../controllers/files';
const router = express.Router();

router.get('/home', controller.getFiles);
router.get('/home/:filename', controller.getFile);
router.post('/home/:filename', controller.updateFile);
router.put('/home/:oldname/:newname', controller.renameFile);
router.delete('/home/:filename', controller.deleteFile);

export = router;
