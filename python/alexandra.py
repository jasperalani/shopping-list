# alexandra.py

# Let's get this party started!
import falcon
import os.path
import mysql.connector
import re
import cgi



# Falcon follows the REST architectural style, meaning (among
# other things) that you think in terms of resources and state
# transitions, which map to HTTP verbs.

# Alexandra
# Image storage, indexing and retreival
class Alexandra(object):

    def on_post(self, req, resp):

        mydb = mysql.connector.connect(
          host="localhost",
          user="root",
          password="password",
          database="alexandra"
        )
        mycursor = mydb.cursor()

        # Create entry in database
        mycursor.execute("INSERT INTO images (id) VALUES (null)")
        mydb.commit()

        # Retreive newly minted id
        mycursor.execute("SELECT MAX(id) FROM images")
        myresult = mycursor.fetchall()

        newlyMintedId = ''

        for x in myresult:
            newlyMintedId = str(x)

        newlyMintedId = ''.join(newlyMintedId)

        newlyMintedId = re.findall(r'\d+', newlyMintedId)[0]

        print(req)
        # TODO: Either validate that content type is multipart/form-data
        #       here, or in another hook before allowing execution to proceed.

        # This must be done to avoid a bug in cgi.FieldStorage
        env = req.env
        env.setdefault('QUERY_STRING', '')

        # # TODO: Add error handling, when the request is not formatted
        # # correctly or does not contain the desired field...

        # # TODO: Consider overriding make_file, so that you can
        # # stream directly to the destination rather than
        # # buffering using TemporaryFile (see http://goo.gl/Yo8h3P)
        print(req.stream)
        # form = cgi.FieldStorage(fp=req.stream, environ=env)
        # print(form)
        # file_item = form[name]
        # if file_item.file:
        #     # TODO: It's an uploaded file... read it in
        #     print(file_item.file)
        #     # with open('picture.png', 'rb') as f:
        #     #     data = f.read()

        #     # with open('picture_out.png', 'wb') as f:
        #     #     f.write(data)
        # else:
        #     quit()
        with open('myfile.txt','w') as f:
            f.write(req.stream) 
        resp.body = '{"newlyMintedId": ' + newlyMintedId + '}'
        resp.status = falcon.HTTP_200

    def on_get(self, req, resp):
        """Handles GET requests"""
        resp.status = falcon.HTTP_200  # This is the default status
        resp.body = ('\nTwo things awe me most, the starry sky '
                     'above me and the moral law within me.\n'
                     '\n'
                     '    ~ Immanuel Kant\n\n')

# falcon.API instances are callable WSGI apps
app = falcon.API()

# Resources are represented by long-lived class instances
alexandra = Alexandra()

# alexandra will handle all requests to the '/alexandra' URL path
app.add_route('/image', alexandra)


def to_utf8(s):
    return s if isinstance(s, str) else s.decode('utf-8')