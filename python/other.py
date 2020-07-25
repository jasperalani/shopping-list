
        # TODO: Either validate that content type is multipart/form-data
        # here, or in another hook before allowing execution to proceed.

        # # This must be done to avoid a bug in cgi.FieldStorage
        # env = req.env
        # env.setdefault('QUERY_STRING', '')

        # # TODO: Add error handling, when the request is not formatted
        # # correctly or does not contain the desired field...

        # # TODO: Consider overriding make_file, so that you can
        # # stream directly to the destination rather than
        # # buffering using TemporaryFile (see http://goo.gl/Yo8h3P)
        # form = cgi.FieldStorage(fp=req.stream, environ=env)

        # file_item = form[name]
        # if file_item.file:
        #     # TODO: It's an uploaded file... read it in

        #     save_path = './images/'

        #     completeName = os.path.join(save_path, name_of_file+".txt")         

        #     file1 = open(completeName, "w")

        #     toFile = raw_input("Write what you want into the field")

        #     file1.write(toFile)

        #     file1.close()
        # else:
        #     # TODO: Raise an error